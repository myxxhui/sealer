// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package local

import (
	"fmt"
	"time"

	"github.com/alibaba/sealer/client/k8s"

	"github.com/alibaba/sealer/image/cache"
	"github.com/pkg/errors"

	"github.com/opencontainers/go-digest"

	"github.com/alibaba/sealer/image/store"

	"github.com/alibaba/sealer/common"
	"github.com/alibaba/sealer/image"
	"github.com/alibaba/sealer/image/reference"
	"github.com/alibaba/sealer/logger"
	"github.com/alibaba/sealer/parser"
	v1 "github.com/alibaba/sealer/types/api/v1"
	"github.com/alibaba/sealer/utils"
)

type builderLayer struct {
	BaseLayers []v1.Layer
	NewLayers  []v1.Layer
}

// LocalBuilder: local builder using local provider to build a cluster image
type Builder struct {
	BuildType        string
	NoCache          bool
	Image            *v1.Image
	Cluster          *v1.Cluster
	ImageNamed       reference.Named
	ImageID          string
	Context          string
	KubeFileName     string
	LayerStore       store.LayerStore
	ImageStore       store.ImageStore
	ImageService     image.Service
	Prober           image.Prober
	FS               store.Backend
	Client           *k8s.Client
	DockerImageCache *MountTarget
	builderLayer
}

func (l *Builder) Build(name string, context string, kubefileName string) error {
	err := l.InitBuilder(name, context, kubefileName)
	if err != nil {
		return err
	}

	k8sClient, err := k8s.Newk8sClient()
	if err != nil {
		return err
	}
	l.Client = k8sClient

	registryCache, err := NewRegistryCache()
	if err != nil {
		return err
	}
	l.DockerImageCache = registryCache
	pipLine, err := l.GetBuildPipeLine()
	if err != nil {
		return err
	}

	for _, f := range pipLine {
		if err = f(); err != nil {
			return err
		}
	}
	return nil
}

func (l *Builder) InitBuilder(name string, context string, kubefileName string) error {
	named, err := reference.ParseToNamed(name)
	if err != nil {
		return err
	}

	absContext, absKubeFile, err := ParseBuildArgs(context, kubefileName)
	if err != nil {
		return err
	}

	err = ValidateContextDirectory(absContext)
	if err != nil {
		return err
	}

	l.ImageNamed = named
	l.Context = absContext
	l.KubeFileName = absKubeFile
	return nil
}

func (l *Builder) GetBuildPipeLine() ([]func() error, error) {
	var buildPipeline []func() error
	if err := l.InitImageSpec(); err != nil {
		return nil, err
	}

	buildPipeline = append(buildPipeline,
		l.PullBaseImageNotExist,
		l.ExecBuild,
		l.CollectRegistryCache,
		l.UpdateImageMetadata,
		l.Cleanup,
	)
	return buildPipeline, nil
}

// init default Image metadata
func (l *Builder) InitImageSpec() error {
	kubeFile, err := utils.ReadAll(l.KubeFileName)
	if err != nil {
		return fmt.Errorf("failed to load kubefile: %v", err)
	}
	l.Image = parser.NewParse().Parse(kubeFile)
	if l.Image == nil {
		return fmt.Errorf("failed to parse kubefile, image is nil")
	}

	layer0 := l.Image.Spec.Layers[0]
	if layer0.Type != FromCmd {
		return fmt.Errorf("first line of kubefile must start with FROM")
	}

	logger.Info("init image spec success!")
	return nil
}

func (l *Builder) PullBaseImageNotExist() (err error) {
	if l.Image.Spec.Layers[0].Value == common.ImageScratch {
		return nil
	}
	if err = l.ImageService.PullIfNotExist(l.Image.Spec.Layers[0].Value); err != nil {
		return fmt.Errorf("failed to pull baseImage: %v", err)
	}
	logger.Info("pull base image %s success", l.Image.Spec.Layers[0].Value)
	return nil
}

func (l *Builder) ExecBuild() error {
	err := l.updateBuilderLayers(l.Image)
	if err != nil {
		return err
	}
	var (
		canUseCache = !l.NoCache
		parentID    = cache.ChainID("")
		newLayers   = l.NewLayers
	)

	baseLayerPaths := GetBaseLayersPath(l.BaseLayers)
	chainSvc, err := cache.NewService()
	if err != nil {
		return err
	}

	hc := handlerContext{
		buildContext:  l.Context,
		continueCache: canUseCache,
		cacheSvc:      chainSvc,
		prober:        l.Prober,
		parentID:      parentID,
		ignoreError:   l.BuildType == common.LiteBuild,
	}

	mhandler := handler{
		hc:         hc,
		layerStore: l.LayerStore,
	}
	for i := 0; i < len(newLayers); i++ {
		// take layer reference
		// we are to modify the layer
		layer := &newLayers[i]
		logger.Info("run build layer: %s %s", layer.Type, layer.Value)
		var (
			layerID digest.Digest
			cacheID digest.Digest
			forErr  error
		)

		switch layer.Type {
		case common.CMDCOMMAND, common.RUNCOMMAND:
			layerID, forErr = mhandler.handleCMDRUNCmd(*layer, baseLayerPaths...)
			if forErr != nil {
				return forErr
			}

		case common.COPYCOMMAND:
			layerID, cacheID, forErr = mhandler.handleCopyCmd(*layer)
			if forErr != nil {
				return forErr
			}
			// hit cache failed, so we save cacheID value to metadata cacheID for this layer.
			// and next time, the cacheID will be used to hit cache.
			if layerID != "" && cacheID != "" && !mhandler.hc.continueCache {
				// TODO set cache id under register.
				forErr = l.SetCacheID(layerID, cacheID.String())
				if forErr != nil {
					logger.Warn("set cache failed layer: %v, err: %v", layer, err)
				}
			}
		}

		layer.ID = layerID
		if layerID == "" {
			continue
		}
		baseLayerPaths = append(baseLayerPaths, l.FS.LayerDataDir(layer.ID))
	}

	logger.Info("exec all build instructs success !")
	return nil
}
func (l *Builder) CollectRegistryCache() error {
	if l.DockerImageCache == nil {
		return nil
	}
	logger.Info("waiting resource to sync")
	//wait resource to sync.do sleep here,because we can't fetch the pod status immediately.
	//if we use retry to check pod status, will pass the cache part, due to some resources has not been created yet.
	time.Sleep(30 * time.Second)
	if !l.IsAllPodsRunning() {
		return fmt.Errorf("cache docker image failed,cluster pod not running")
	}
	imageLayer := v1.Layer{
		Type:  imageLayerType,
		Value: "registry cache",
	}
	layerDgst, err := l.RegisterLayer(l.DockerImageCache.GetMountUpper())
	if err != nil {
		return err
	}

	imageLayer.ID = layerDgst
	l.NewLayers = append(l.NewLayers, imageLayer)

	logger.Info("save image cache success")
	return nil
}

//This function only has meaning for copy layers
func (l *Builder) SetCacheID(layerID digest.Digest, cID string) error {
	return l.FS.SetMetadata(layerID, cacheID, []byte(cID))
}

func (l *Builder) RegisterLayer(tempTarget string) (digest.Digest, error) {
	layerDigest, err := l.LayerStore.RegisterLayerForBuilder(tempTarget)
	if err != nil {
		return "", fmt.Errorf("failed to register layer, err: %v", err)
	}

	return layerDigest, nil
}

func (l *Builder) UpdateImageMetadata() error {
	err := setClusterFileToImage(l.Image, l.ImageNamed.Raw())
	if err != nil {
		return fmt.Errorf("failed to set image metadata, err: %v", err)
	}

	l.Image.Spec.Layers = append(l.BaseLayers, l.NewLayers...)

	err = l.updateImageIDAndSaveImage()
	if err != nil {
		return fmt.Errorf("failed to save image metadata, err: %v", err)
	}

	logger.Info("update image %s to image metadata success !", l.ImageNamed.Raw())
	return nil
}

func (l *Builder) updateImageIDAndSaveImage() error {
	imageID, err := generateImageID(*l.Image)
	if err != nil {
		return err
	}

	l.Image.Spec.ID = imageID
	return l.ImageStore.Save(*l.Image, l.ImageNamed.Raw())
}

func (l *Builder) updateBuilderLayers(image *v1.Image) error {
	// we do not check the len of layers here, because we checked it before.
	// remove the first layer of image
	var (
		layer0    = image.Spec.Layers[0]
		baseImage *v1.Image
		err       error
	)

	// and the layer 0 must be from layer
	if layer0.Value == common.ImageScratch {
		// give an empty image
		baseImage = &v1.Image{}
	} else {
		baseImage, err = l.ImageStore.GetByName(image.Spec.Layers[0].Value)
		if err != nil {
			return fmt.Errorf("failed to get base image while updating base layers, err: %s", err)
		}
	}

	l.BaseLayers = append([]v1.Layer{}, baseImage.Spec.Layers...)
	l.NewLayers = append([]v1.Layer{}, image.Spec.Layers[1:]...)
	if len(l.BaseLayers)+len(l.NewLayers) > maxLayerDeep {
		return errors.New("current number of layers exceeds 128 layers")
	}
	return nil
}
func (l *Builder) Cleanup() (err error) {
	// umount registry
	if l.DockerImageCache != nil {
		l.DockerImageCache.CleanUp()
		return
	}

	return err
}

func (l *Builder) IsAllPodsRunning() bool {
	err := utils.Retry(10, 5*time.Second, func() error {
		namespacePodList, err := l.Client.ListAllNamespacesPods()
		if err != nil {
			return err
		}

		var notRunning int
		for _, podNamespace := range namespacePodList {
			for _, pod := range podNamespace.PodList.Items {
				if pod.Status.Phase != "Running" && pod.Status.Phase != "Succeeded" {
					logger.Info(podNamespace.Namespace.Name, pod.Name, pod.Status.Phase)
					notRunning++
					continue
				}
			}
		}
		if notRunning > 0 {
			logger.Info("remaining %d pod not running", notRunning)
			return fmt.Errorf("pod not running")
		}
		return nil
	})
	return err == nil
}
