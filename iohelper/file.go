// Package iohelper
// Date: 2022/9/7 00:43
// Author: Amu
// Description:
package iohelper

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

// FileExist 检查文件是否存在
func FileExist(absPath string) bool {
	exist := true
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// DeleteFile 删除文件
func DeleteFile(absPath string) error {
	return os.RemoveAll(absPath)
}

// GetPathFiles 获取指定目录下的所有文件
func GetPathFiles(absDir string) (re []string) {
	if PathExist(absDir) {
		files, _ := os.ReadDir(absDir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetPwd 获取程序运行目录
func GetPwd() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

// FileMD5 计算文件的 md5
func FileMD5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5Handle := md5.New()
	_, err = io.Copy(md5Handle, f)
	if err != nil {
		return "", err
	}
	md5String := hex.EncodeToString(md5Handle.Sum(nil))
	return md5String, nil
}

// GetFileSize 获取文件大小
func GetFileSize(filename string) (int64, error) {
	file, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return file.Size(), nil
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
	if !PathExist(des) {
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
