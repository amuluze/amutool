// Package file
// Date: 2022/9/7 00:43
// Author: Amu
// Description:
package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// CheckPathIsExist 检查文件或目录是否存在
func CheckPathIsExist(absPath string) bool {
	exist := true
	if _, err := os.Stat(absPath); os.IsExist(err) {
		exist = false
	}
	return exist
}

// MakeDir 创建目录
func MakeDir(absDir string) error {
	return os.MkdirAll(path.Dir(absDir), os.ModePerm)
}

// Delete 删除文件或目录
func Delete(absPath string) error {
	return os.RemoveAll(absPath)
}

// GetPathDirs 获取指定目录下的所有文件夹
func GetPathDirs(absDir string) (re []string) {
	if CheckPathIsExist(absDir) {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetPathFiles 获取指定目录下的所有文件
func GetPathFiles(absDir string) (re []string) {
	if CheckPathIsExist(absDir) {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetModelPath 获取程序运行目录
func GetModelPath() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

// SaveToFile 写入文件
func SaveToFile(fname string, src []string, isClear bool) bool {
	return WriteFile(fname, src, isClear)
}

// WriteFile 写入文件
func WriteFile(fname string, src []string, isClear bool) bool {
	MakeDir(fname)
	flag := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !isClear {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	}
	f, err := os.OpenFile(fname, flag, 0666)
	if err != nil {
		return false
	}
	defer f.Close()

	for _, v := range src {
		f.WriteString(v)
		f.WriteString("\r\n")
	}

	return true
}

// ReadFile 读取文件
func ReadFile(fname string) (src []string) {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		src = append(src, string(line))
	}

	return src
}

// MoveFile 移动文件或文件夹(/结尾)
func MoveFile(from, to string) error {
	// if !CheckFileIsExist(to) {
	// 	BuildDir(to)
	// }
	return os.Rename(from, to)
}

func CopyFile(src, des string) error {
	if !CheckPathIsExist(des) {
		MakeDir(des)
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		return err
	}
	defer desFile.Close()

	_, err = io.Copy(desFile, srcFile)
	return err
}
