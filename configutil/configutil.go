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
	var configFilePath = "config.yaml"
	configFile, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		logutil.Fatal("Read config yaml file err %v" + err.Error())
	}
}

// 将配置文件参数转换为模型
//
//	配置文件为根目录下config.yaml
//	modelReffrence:模型的地址
func ToModel(modelReffrence interface{}) error {
	return yaml.Unmarshal(configFile, modelReffrence)
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
