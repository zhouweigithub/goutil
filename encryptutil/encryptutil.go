package encryptutil

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// md5加密
func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	re := h.Sum(nil)
	return fmt.Sprintf("%x", re)
}

// base64标准加密
func Base64Encodeing(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64标准解密
func Base64Decodeing(str string) (string, error) {
	var bytes, err = base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
