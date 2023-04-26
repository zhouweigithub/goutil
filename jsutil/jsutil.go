package jsutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dop251/goja"
)

/*
js操作类

用法：先写一个 struct 继承此类，再编写需要映射的方法

如：

写一个 struct 继承此类

	type jsMap struct {
		jsutil.JsHelper
	}

编写需要映射的方法

	func (j *jsMap) AddInt() (fn func(int, int) int, err error) {
		var tmp = j.Vm.Get("addInt")
		if tmp == nil {
			err = errors.New("Js函数 addInt 映射到 Go 函数失败！js中未找到该函数 addInt")
		} else {
			err = j.Vm.ExportTo(tmp, &fn)
			if err != nil {
				err = errors.New("Js函数 addInt 映射到 Go 函数失败！\n" + err.Error())
			}
		}
		return
	}

调用示例：

	var js jsMap
	if err := js.LoadFile("test.js"); err != nil {
		fmt.Println(err.Error())
	} else {
		var a, err = js.AddInt()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(a(3, 5))
		}
	}
*/
type JsHelper struct {
	Vm *goja.Runtime
}

// 初始化并加载js文件
//
//	jsFilePaths: js文件地址
func (j *JsHelper) LoadFile(jsFilePaths ...string) error {
	if err := j.init(jsFilePaths...); err != nil {
		return err
	}
	return nil
}

// 读取js文件
func (j *JsHelper) loadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 初始化
//
//	jsFilePaths: js文件地址
func (j *JsHelper) init(jsFilePaths ...string) error {
	if len(jsFilePaths) == 0 {
		return errors.New("jsHelper初始化失败，未提供js文件路径")
	}
	j.Vm = goja.New()
	var sb strings.Builder
	for i := range jsFilePaths {
		var content, err = j.loadFile(jsFilePaths[i])
		if err != nil {
			return err
		}
		sb.WriteString(content + ";")
	}
	_, err := j.Vm.RunString(sb.String())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
