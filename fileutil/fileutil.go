package fileutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
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

// 覆盖已有文件，文件不存在则创建，目录需要提前创建
func WriteTextFile(path string, content string) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		log.Println("flush error :", err)
		return err
	}
	return nil
}

// 追加内容到文本文件末尾，文件不存在则创建，目录需要提前创建
func AppendTextFile(path string, content string) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		log.Println("flush error :", err)
		return err
	}
	return nil
}

// 如果不存在就创建目录
func CreateFolderIfNotExists(folder string) {
	if !IsExists(folder) {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
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

// 获取文件路径的目录路径
func GetFolder(filePath string) string {
	return path.Dir(filePath)
}

// 获取目录中满足正则式的所有目录
func GetFolders(folder string, isLoop bool) []string {
	files, _ := ioutil.ReadDir(folder)
	var filePaths = []string{}
	for _, file := range files {
		if file.IsDir() {
			filePaths = append(filePaths, file.Name())
			if isLoop {
				filePaths = append(filePaths, GetFiles(folder+"/"+file.Name(), true)...)
			}
		}
	}
	return filePaths
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

// 删除文件或者目录
func DeleteFileOrFolder(path string) error {
	err := os.Remove(path)
	return err
}

// 复制文件
//
//	fromFile：from file
//	toFile：to file, if folder not exists create it
//	return copyed bytes count, copyed error
func CopyFile(fromFile, toFile string) (int64, error) {
	sourceFileStat, err := os.Stat(fromFile)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", fromFile)
	}

	source, err := os.Open(fromFile)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	var toFolder = path.Dir(toFile)
	CreateFolderIfNotExists(toFolder)
	destination, err := os.Create(toFile)
	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// 复制文件夹
//
//	fromFolder：from folder
//	toFolder：to folder, if not exists create it
func CopyFolder(fromFolder, toFolder string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(fromFolder); err != nil {
		return err
	}
	if err = os.MkdirAll(toFolder, srcinfo.Mode()); err != nil {
		return err
	}
	if fds, err = ioutil.ReadDir(fromFolder); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(fromFolder, fd.Name())
		dstfp := path.Join(toFolder, fd.Name())

		if fd.IsDir() {
			if err = CopyFolder(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if _, err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
