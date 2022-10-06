/*
@Time : 2022/10/2 17:33
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : directory
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package utils

import (
	"garavel/global"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func RootPath() string {
	var rootDir string
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	rootDir = filepath.Dir(filepath.Dir(exePath))
	tmpDir := os.TempDir()
	if strings.Contains(exePath, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			rootDir = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
		}
	}
	return filepath.Join(rootDir, "..")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.G_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.G_LOG.Error("create directory"+ v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
