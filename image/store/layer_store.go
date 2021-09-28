package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/alibaba/sealer/common"
	"github.com/alibaba/sealer/logger"
	"github.com/alibaba/sealer/utils"
	"github.com/opencontainers/go-digest"
)

type layerStore struct {
	mux    sync.RWMutex
	layers map[LayerID]*roLayer
}

func (ls *layerStore) Get(id LayerID) Layer {
	ls.mux.RLock()
	defer ls.mux.RUnlock()
	l, ok := ls.layers[id]
	if !ok {
		return nil
	}
	return l
}

func (ls *layerStore) RegisterLayerIfNotPresent(layer Layer) error {
	layerExist := ls.Get(layer.ID())
	if layerExist != nil {
		return nil
	}

	err := dumpLayerMetadata(layer)
	if err != nil {
		return err
	}

	ls.mux.Lock()
	defer ls.mux.Unlock()
	if roLayer, ok := layer.(*roLayer); ok {
		ls.layers[layer.ID()] = roLayer
	}

	return nil
}

func (ls *layerStore) Delete(id LayerID) error {
	digs := digest.Digest(id)
	layer := ls.Get(id)
	if layer == nil {
		logger.Debug("layer %s is already deleted", id)
		return nil
	}

	filePath := filepath.Join(common.DefaultLayerDir, digs.Hex())
	err := os.RemoveAll(filePath)
	if err != nil {
		return err
	}
	subDir := filepath.Join(common.DefaultLayerDBDir, digs.Algorithm().String(), digs.Hex())
	err = os.RemoveAll(subDir)
	if err != nil {
		return err
	}
	ls.mux.Lock()
	defer ls.mux.Unlock()
	delete(ls.layers, id)
	return nil
}

func dumpLayerMetadata(layer Layer) error {
	id := layer.ID()
	digs := digest.Digest(id)
	subDir := filepath.Join(common.DefaultLayerDBDir, digs.Algorithm().String(), digs.Hex())
	return utils.WriteFile(filepath.Join(subDir, "id"), []byte(id.String()))
}

func getDirListInDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, filepath.Join(dir, file.Name()))
		}
	}
	return dirs, nil
}

func getAllROLayers() ([]*roLayer, error) {
	err := utils.MkDirIfNotExists(common.DefaultLayerDBDir)
	if err != nil {
		return nil, err
	}
	// TODO maybe there no need to traverse layerdb, just clarify how may sha supported in a list
	shaDirs, err := getDirListInDir(common.DefaultLayerDBDir)
	if err != nil {
		return nil, err
	}

	var layerDirs []string
	for _, shaDir := range shaDirs {
		layerDirList, err := getDirListInDir(shaDir)
		if err != nil {
			return nil, err
		}
		layerDirs = append(layerDirs, layerDirList...)
	}

	var res []*roLayer
	for _, layerDir := range layerDirs {
		id, err := ioutil.ReadFile(filepath.Join(layerDir, "id"))
		if err == nil {
			_, err := digest.Parse(string(id))
			if err == nil {
				res = append(res, &roLayer{id: LayerID(id)})
			} else {
				logger.Warn("failed to get layer metadata %s, which has a invalid id, err: %s", filepath.Base(layerDir), err)
			}
		} else {
			logger.Warn("failed to get layer metadata %s, whose id file lost, err: %s", filepath.Base(layerDir), err)
		}
	}

	return res, nil
}

func NewDefaultLayerStore() (LayerStore, error) {
	ls := &layerStore{layers: map[LayerID]*roLayer{}}
	layers, err := getAllROLayers()
	if err != nil {
		return nil, err
	}

	ls.mux.Lock()
	defer ls.mux.Unlock()
	//TODO only check .../layerdb/.../id for existence of layer currently
	for _, layer := range layers {
		ls.layers[layer.id] = layer
	}

	return ls, nil
}
