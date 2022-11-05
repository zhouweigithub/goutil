package configutil

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/zhouweigithub/goutil/logutil"

	"gopkg.in/yaml.v2"
)

var configFile []byte

func init() {
	var err error
	var configFilePath = filepath.Join(getCurrentAbPathByCaller(), "config.yaml")
	configFile, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		logutil.Fatal("Read config yaml file err %v" + err.Error())
	}
}

func ToModel() (e interface{}, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}

// 获取程序运行路径（go build）
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logutil.Error("Get current path err %v" + err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
