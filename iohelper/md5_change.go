// Package iohelper
// Date: 2023/3/28 17:14
// Author: Amu
// Description:
package iohelper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type CheckTextMd5 struct {
	fileName    string
	fileText    string
	versionFile string
}

func textMD5(s string) string {
	md5Handle := md5.New()
	_, _ = io.WriteString(md5Handle, s)
	return hex.EncodeToString(md5Handle.Sum(nil))
}

// NewCheckTextMd5 fileName 和 fileText 人选一个传入， md5Name 是 md5 文件名称
func NewCheckTextMd5(fileName, fileText, md5FilePath, md5Name string) *CheckTextMd5 {
	if fileName == "" && fileText == "" {
		panic("NewCheckTextMd5 fail, fileName fileText both null")
	}

	var namePrefix string
	if fileText != "" {
		namePrefix = textMD5(fileText)
	} else {
		namePrefix, _ = FileMD5(fileName)
	}

	versionFile := fmt.Sprint(md5FilePath, "/.", namePrefix, "-", md5Name)
	fmt.Printf("version file: %v\n", versionFile)
	err := EnsureDir(versionFile)
	if err != nil {
		panic("create template md5 file dir failure")
	}

	return &CheckTextMd5{
		fileName:    fileName,
		fileText:    fileText,
		versionFile: versionFile,
	}
	return nil
}

func (s *CheckTextMd5) Change() bool {
	var err error

	versionExist := FileExist(s.versionFile)
	var templateFileMd5 string
	var newTemplateFileMd5 string

	templateFileChanged := false
	if versionExist {
		templateFIleByte, _ := os.ReadFile(s.versionFile)
		templateFileMd5 = string(templateFIleByte)
		if s.fileText != "" {
			newTemplateFileMd5 = textMD5(s.fileText)
		} else {
			newTemplateFileMd5, _ = FileMD5(s.fileName)
		}

		if newTemplateFileMd5 != templateFileMd5 && newTemplateFileMd5 != "" {
			templateFileChanged = true
		}
	} else {
		templateFileChanged = true
		err = s.Write()
	}

	if err != nil {
		panic("template change md5 get failure")
	}
	return templateFileChanged
}

func (s *CheckTextMd5) Write() error {
	var templateFileMd5 string
	if s.fileText != "" {
		templateFileMd5 = textMD5(s.fileText)
	} else {
		templateFileMd5, _ = FileMD5(s.fileName)
	}
	fmt.Printf("template file md5: %v\n", templateFileMd5)
	return os.WriteFile(s.versionFile, []byte(templateFileMd5), 0640)
}
