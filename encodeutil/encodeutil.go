package encodeutil

import (
	"bytes"
	"io/ioutil"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 转换中文，防止乱码
func EncodeHtmlBytes(html []byte) []byte {
	if !utf8.Valid(html) {
		bt, _ := simplifiedchinese.GBK.NewDecoder().Bytes(html)
		return bt
	}
	return html
}

// 转换中文，防止乱码
func EncodeHtmlString(html string) string {
	var bytes = []byte(html)
	if !utf8.Valid(bytes) {
		bt, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)
		return string(bt)
	}
	return html
}
