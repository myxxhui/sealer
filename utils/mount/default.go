package mount

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type Default struct {
}

// Unmount target
func (d *Default) Unmount(target string) error {
	err := os.RemoveAll(target)
	if err != nil {
		return fmt.Errorf("remote target failed: %s", err)
	}
	return nil
}

// copy all layers to target merged dir
func (d *Default) Mount(target string, upperDir string, layers ...string) error {
	//if target is empty,return err
	if target == "" {
		return fmt.Errorf("target is empty")
	}

	reverse(layers)

	for _, layer := range layers {
		srcInfo, err := os.Stat(layer)
		if err != nil {
			return fmt.Errorf("get srcInfo err: %s", err)
		}
		if srcInfo.IsDir() {
			err := copyDir(layer, target)
			if err != nil {
				return fmt.Errorf("copyDir [%s] to [%s] failed: %s", layer, target, err)
			}
		} else {
			IsExist, err := PathExists(target)
			if err != nil {
				return err
			}
			if !IsExist {
				err = os.Mkdir(target, 0666)
				if err != nil {
					return fmt.Errorf("mkdir [%s] error %v", target, err)
				}
			}
			_file := filepath.Base(layer)
			dst := path.Join(target, _file)
			err = copyFile(layer, dst)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func copyDir(srcPath string, dstPath string) error {
	IsExist, err := PathExists(dstPath)
	if err != nil {
		return err
	}
	if !IsExist {
		err = os.Mkdir(dstPath, 0666)
		if err != nil {
			return fmt.Errorf("mkdir [%s] error %v", dstPath, err)
		}
	}

	srcFiles, err := ioutil.ReadDir(srcPath)
	if err != nil {
		return err
	}
	for _, file := range srcFiles {
		src := path.Join(srcPath, file.Name())
		dst := path.Join(dstPath, file.Name())
		if file.IsDir() {
			err = copyDir(src, dst)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(src, dst)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	// open srcfile
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open file [%s] failed: %s", src, err)
	}
	defer srcFile.Close()

	// create dstfile
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create file err: %s", err)
	}
	defer dstFile.Close()
	// copy  file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("copy file err: %s", err)
	}
	return nil
}

//notExist false ,Exist true
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("os.Stat(%s) err: %s", path, err)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
