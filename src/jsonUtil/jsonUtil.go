package jsonutil

import (
	"encoding/json"
)

func ToJson(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	} else {
		str := string(bytes)
		return str, nil
	}
}

func ToModel(jsonString string, model interface{}) error {
	err := json.Unmarshal([]byte(jsonString), &model)
	if err != nil {
		return err
	} else {
		return nil
	}
}
