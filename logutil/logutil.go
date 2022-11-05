package logutil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var logFileFolder string
var mutex sync.RWMutex //读写锁
type logType int

const (
	info logType = iota
	debug
	warm
	error
	fatal
)

func SetLogFolder(folder string) {
	logFileFolder = folder
}
func Info(msg string) {
	write(msg, info)
}
func Debug(msg string) {
	write(msg, debug)
}
func Warm(msg string) {
	write(msg, warm)
}
func Error(msg string) {
	write(msg, error)
}
func Fatal(msg string) {
	write(msg, fatal)
}
func write(msg string, logType logType) {
	//mutex.Lock()	//锁住写操作
	//defer  mutex.Unlock()  //方法退出前执行

	if logFileFolder == "" {
		logFileFolder = "Log"
	}

	var now = time.Now()
	var year = now.Year()
	var month = fmt.Sprintf("%d", time.Now().Month())

	var logFolder = logFileFolder
	logFolder = logFolder + "/" + strconv.Itoa(year) + "/" + month

	isExist := pathExists(logFolder)
	if !isExist {
		err1 := os.MkdirAll(logFolder, os.ModePerm)
		if err1 != nil {
			log.Println(err1)
			return
		}
	}

	file, err := os.OpenFile(getLogFileName(logType, logFolder), os.O_APPEND|os.O_CREATE, 0644)
	if err == nil {
		defer file.Close()
		_, err2 := file.WriteString(formatMsg(msg))
		if err2 != nil {
			log.Println(err.Error())
			return
		}
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
		//log.Println(err)
	}
	return false
}

func formatMsg(msg string) string {
	var timeString = time.Now().Format("2006-01-02 15:04:05")
	return "#" + timeString + " " + msg + "\r\n" + "------------------------------------------------------\r\n"
}

func getFileNameSuffix() string {
	return time.Now().Format("2006-01-02.15")
}

func getLogFileName(logType logType, logPath string) string {
	var typeText string
	switch logType {
	case info:
		typeText = "Info"
	case debug:
		typeText = "Debug"
	case warm:
		typeText = "Warm"
	case error:
		typeText = "Error"
	case fatal:
		typeText = "Fatal"
	}

	return logPath + "\\" + getFileNameSuffix() + "." + typeText + ".txt"
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
