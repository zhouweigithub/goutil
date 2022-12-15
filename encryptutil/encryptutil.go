package encryptutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
)

// md5加密
func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
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

// hmacsha256验证
func HMAC_SHA256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// hmacsha512验证
func HMAC_SHA512(src, key string) string {
	m := hmac.New(sha512.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// hmacsha1验证
func HMAC_SHA1(src, key string) string {
	m := hmac.New(sha1.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// sha256验证
func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// sha512验证
func SHA512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

var ivspec = []byte("0000000000000000")

// AES，即高级加密标准（Advanced Encryption Standard），是一个对称分组密码算法，旨在取代DES成为广泛使用的标准。
//
// AES中常见的有三种解决方案，分别为AES-128、AES-192和AES-256。
//
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func AESEncodeStr(src, key string) (string, error) {
	if src == "" {
		return "", errors.New("plain content empty")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	ecb := cipher.NewCBCEncrypter(block, ivspec)
	content := []byte(src)
	content = _PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return hex.EncodeToString(crypted), nil
}

// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func AESDecodeStr(crypt, key string) (string, error) {
	if crypt == "" {
		return "", errors.New("plain content empty")
	}
	crypted, err := hex.DecodeString(strings.ToLower(crypt))
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	ecb := cipher.NewCBCDecrypter(block, ivspec)
	decrypted := make([]byte, len(crypted))
	ecb.CryptBlocks(decrypted, crypted)

	return string(_PKCS5Trimming(decrypted)), nil
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// 中文转Unicode -> \u652f
func ToUnicode(text string) string {
	textQuoted := strconv.QuoteToASCII(text)
	return textQuoted[1 : len(textQuoted)-1]
}

// Unicode转中文 <- \u652f
func FromUnicode(text string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(text)), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}
