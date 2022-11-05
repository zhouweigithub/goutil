package fileutil

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/zhouweigithub/goutil/logutil"
)

// 读取文件所有内容字符串
func GetFileContent(path string) string {
	res, err := ioutil.ReadFile(path)
	if err != nil {
		errMsg := "读取文件出错！" + path + "\r\n" + err.Error()
		log.Println(errMsg)
		logutil.Error(errMsg)
		return ""
	} else {
		return string(res)
	}
}

// 读取文件各行数据
func GetFileLines(path string) []string {
	content := GetFileContent(path)
	if content == "" {
		return nil
	} else {
		return strings.Split(content, "\r\n")
	}
}

// 写入文本文件
func WriteTextFile(path string, content string) bool {
	if !IsExists(path) {
		os.Create(path)
	}
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		errMsg := "写入文件出错！" + path + "\r\n" + err.Error()
		log.Println(errMsg)
		logutil.Error(errMsg)
		return false
	} else {
		return true
	}
}

// 如果不存在就创建目录
func CreateFolderIfNotExists(path string) {
	if !IsExists(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Println(err.Error())
			logutil.Error(err.Error())
		}
	}
}

// 检测文件或文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取子目录所有文件（仅文件名）
func GetFiles(folder string, isLoop bool) []string {
	files, _ := ioutil.ReadDir(folder)
	var filePaths = []string{}
	for _, file := range files {
		if file.IsDir() {
			if isLoop {
				filePaths = append(filePaths, GetFiles(folder+"/"+file.Name(), true)...)
			}
		} else {
			filePaths = append(filePaths, file.Name())
		}
	}
	return filePaths
}

// 获取子目录所有文件（包含相对路径）
func GetFilesWithRelativePath(folder string, isLoop bool) []string {
	return getFilesWithRelativePaths(folder, folder, isLoop)
}

// 获取子目录所有文件（包含相对路径）
func getFilesWithRelativePaths(folder string, baseFolder string, isLoop bool) []string {
	files, _ := ioutil.ReadDir(folder)
	var filePaths = []string{}
	for _, file := range files {
		if file.IsDir() {
			if isLoop {
				filePaths = append(filePaths, getFilesWithRelativePaths(folder+"\\"+file.Name(), baseFolder, true)...)
			}
		} else {
			filePaths = append(filePaths, folder+"\\"+file.Name())
		}
	}
	return filePaths
}
