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

package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/alibaba/sealer/image/store"

	"github.com/alibaba/sealer/runtime"

	"github.com/alibaba/sealer/utils"

	"github.com/pkg/errors"

	"github.com/alibaba/sealer/logger"

	"github.com/alibaba/sealer/common"
	"github.com/alibaba/sealer/image"

	v1 "github.com/alibaba/sealer/types/api/v1"
	"github.com/alibaba/sealer/utils/mount"
	"github.com/alibaba/sealer/utils/ssh"
)

const (
	RemoteChmod = "cd %s  && chmod +x scripts/* && cd scripts && bash init.sh"
)

type Interface interface {
	MountRootfs(cluster *v1.Cluster, hosts []string) error
	UnMountRootfs(cluster *v1.Cluster) error
	MountImage(cluster *v1.Cluster) error
	UnMountImage(cluster *v1.Cluster) error
	Clean(cluster *v1.Cluster) error
}

type FileSystem struct {
	imageStore store.ImageStore
}

func (c *FileSystem) Clean(cluster *v1.Cluster) error {
	return utils.CleanFiles(common.GetClusterWorkDir(cluster.Name), common.DefaultClusterBaseDir(cluster.Name), common.DefaultKubeConfigDir())
}

func (c *FileSystem) umountImage(cluster *v1.Cluster) error {
	mountdir := common.DefaultMountCloudImageDir(cluster.Name)
	if utils.IsFileExist(mountdir) {
		var err error
		err = utils.Retry(10, time.Second, func() error {
			err = mount.NewMountDriver().Unmount(mountdir)
			if err != nil {
				return err
			}
			return os.RemoveAll(mountdir)
		})
		if err != nil {
			logger.Warn("failed to unmount dir %s,err: %v", mountdir, err)
		}
	}
	return nil
}

func (c *FileSystem) mountImage(cluster *v1.Cluster) error {
	var (
		mountdir = common.DefaultMountCloudImageDir(cluster.Name)
		upperDir = filepath.Join(mountdir, "upper")
		driver   = mount.NewMountDriver()
		err      error
	)
	if isMount, _ := mount.GetMountDetails(mountdir); isMount {
		err = driver.Unmount(mountdir)
		if err != nil {
			return fmt.Errorf("%s already mount, and failed to umount %v", mountdir, err)
		}
	}
	if utils.IsFileExist(mountdir) {
		err = os.RemoveAll(mountdir)
		if err != nil {
			return fmt.Errorf("failed to clean %s, %v", mountdir, err)
		}
	}
	//get layers
	Image, err := c.imageStore.GetByName(cluster.Spec.Image)
	if err != nil {
		return err
	}
	layers, err := image.GetImageLayerDirs(Image)
	if err != nil {
		return fmt.Errorf("get layers failed: %v", err)
	}

	if err = os.MkdirAll(upperDir, 0744); err != nil {
		return fmt.Errorf("create upperdir failed, %s", err)
	}
	if err = driver.Mount(mountdir, upperDir, layers...); err != nil {
		return fmt.Errorf("mount files failed %v", err)
	}
	return nil
}

func (c *FileSystem) MountImage(cluster *v1.Cluster) error {
	return c.mountImage(cluster)
}

func (c *FileSystem) UnMountImage(cluster *v1.Cluster) error {
	return c.umountImage(cluster)
}

func (c *FileSystem) MountRootfs(cluster *v1.Cluster, hosts []string) error {
	clusterRootfsDir := common.DefaultTheClusterRootfsDir(cluster.Name)
	//scp roofs to all Masters and Nodes,then do init.sh
	if err := mountRootfs(hosts, clusterRootfsDir, cluster); err != nil {
		return fmt.Errorf("mount rootfs failed %v", err)
	}
	return nil
}

func (c *FileSystem) UnMountRootfs(cluster *v1.Cluster) error {
	//do clean.sh,then remove all Masters and Nodes roofs
	IPList := append(cluster.Spec.Masters.IPList, cluster.Spec.Nodes.IPList...)
	config := runtime.GetRegistryConfig(common.DefaultTheClusterRootfsDir(cluster.Name), cluster.Spec.Masters.IPList[0])
	if utils.NotIn(config.IP, IPList) {
		IPList = append(IPList, config.IP)
	}
	if err := unmountRootfs(IPList, cluster); err != nil {
		return err
	}
	return nil
}

func mountRootfs(ipList []string, target string, cluster *v1.Cluster) error {
	SSH := ssh.NewSSHByCluster(cluster)
	config := runtime.GetRegistryConfig(
		common.DefaultTheClusterRootfsDir(cluster.Name),
		cluster.Spec.Masters.IPList[0])
	if err := ssh.WaitSSHReady(SSH, 4, ipList...); err != nil {
		return errors.Wrap(err, "check for node ssh service time out")
	}
	var wg sync.WaitGroup
	var flag bool
	var mutex sync.Mutex
	src := common.DefaultMountCloudImageDir(cluster.Name)
	// TODO scp sdk has change file mod bug
	initCmd := fmt.Sprintf(RemoteChmod, target)
	for _, ip := range ipList {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			err := CopyFiles(SSH, ip == config.IP, ip, src, target)
			if err != nil {
				logger.Error("copy rootfs failed %v", err)
				mutex.Lock()
				flag = true
				mutex.Unlock()
			}
			err = SSH.CmdAsync(ip, initCmd)
			if err != nil {
				logger.Error("exec init.sh failed %v", err)
				mutex.Lock()
				flag = true
				mutex.Unlock()
			}
		}(ip)
	}
	wg.Wait()
	if flag {
		return fmt.Errorf("mountRootfs failed")
	}
	return nil
}

func CopyFiles(ssh ssh.Interface, isRegistry bool, ip, src, target string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to copy files %s", err)
	}

	if isRegistry {
		return ssh.Copy(ip, src, target)
	}
	for _, f := range files {
		if f.Name() == common.RegistryDirName {
			continue
		}
		err = ssh.Copy(ip, filepath.Join(src, f.Name()), filepath.Join(target, f.Name()))
		if err != nil {
			return fmt.Errorf("failed to copy sub files %v", err)
		}
	}
	return nil
}

func unmountRootfs(ipList []string, cluster *v1.Cluster) error {
	SSH := ssh.NewSSHByCluster(cluster)
	var wg sync.WaitGroup
	var flag bool
	var mutex sync.Mutex
	clusterRootfsDir := common.DefaultTheClusterRootfsDir(cluster.Name)
	execClean := fmt.Sprintf("/bin/bash -c "+common.DefaultClusterClearBashFile, cluster.Name)
	rmRootfs := fmt.Sprintf("rm -rf %s", clusterRootfsDir)
	for _, ip := range ipList {
		wg.Add(1)
		go func(IP string) {
			defer wg.Done()
			cmd := fmt.Sprintf("%s && %s", execClean, rmRootfs)
			if mount, _ := mount.GetRemoteMountDetails(SSH, IP, clusterRootfsDir); mount {
				cmd = fmt.Sprintf("umount %s && %s", clusterRootfsDir, cmd)
			}
			if err := SSH.CmdAsync(IP, cmd); err != nil {
				logger.Error("%s:exec %s failed, %s", IP, execClean, err)
				mutex.Lock()
				flag = true
				mutex.Unlock()
				return
			}
		}(ip)
	}
	wg.Wait()
	if flag {
		return fmt.Errorf("unmountRootfs failed")
	}
	return nil
}

func NewFilesystem() (Interface, error) {
	dis, err := store.NewDefaultImageStore()
	if err != nil {
		return nil, err
	}

	return &FileSystem{imageStore: dis}, nil
}
