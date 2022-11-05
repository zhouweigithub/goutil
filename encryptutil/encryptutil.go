package encryptutil

import (
	"crypto/md5"
	"fmt"
)

// md5加密
func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	re := h.Sum(nil)
	return fmt.Sprintf("%x", re)
}
