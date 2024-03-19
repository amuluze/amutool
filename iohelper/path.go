// Package iohelper
// Date: 2023/3/28 17:13
// Author: Amu
// Description:
package iohelper

import (
	"os"
	"path"
)

// PathExist 检查目录是否存在
func PathExist(dir string) bool {
	exist := true
	if _, err := os.Stat(dir); os.IsExist(err) {
		exist = false
	}
	return exist
}

// MakeDir 创建目录
func MakeDir(absDir string) error {
	return os.MkdirAll(path.Dir(absDir), os.ModePerm)
}

// EnsureDir 确保目录存在，如果不存在则创建
func EnsureDir(dir string) error {
	if err := PathExist(dir); !err {
		err := MakeDir(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPathDirs 获取指定目录下的所有文件夹
func GetPathDirs(absDir string) (re []string) {
	if PathExist(absDir) {
		files, _ := os.ReadDir(absDir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}
