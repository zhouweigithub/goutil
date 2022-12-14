package jsonutil

import (
	"encoding/json"
)

// 模型转换为JSON字符串
func ToJson(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	} else {
		str := string(bytes)
		return str, nil
	}
}

// JSON字符串转换为模型
func ToModel(jsonString string, model interface{}) error {
	err := json.Unmarshal([]byte(jsonString), &model)
	if err != nil {
		return err
	} else {
		return nil
	}
}
