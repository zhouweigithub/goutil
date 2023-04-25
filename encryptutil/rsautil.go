package encryptutil

import (
	"encoding/hex"
	"strings"

	"github.com/wenzhenxi/gorsa"
)

type rsa struct {
}

func Rsa() rsa {
	return rsa{}
}

const (
	// 公钥前缀
	publicPrefix = "-----BEGIN Public key-----"
	// 公钥后缀
	publicSuffix = "-----END Public key-----"
	// 私钥前缀
	privatePrefix = "-----BEGIN Private key-----"
	// 私钥后缀
	privateSuffix = "-----END Private key-----"
)

// 公钥加密
func (s rsa) PublicEncrypt(data, publicKey string) (string, error) {
	publicKey = validPublic(publicKey)
	return gorsa.PublicEncrypt(data, publicKey)
}

// 公钥解密
func (s rsa) PublicDecrypt(data, publicKey string) (string, error) {
	publicKey = validPublic(publicKey)
	if d, err := gorsa.PublicDecrypt(data, publicKey); err != nil {
		return "", err
	} else {
		h, _ := hex.DecodeString(d)
		return string(h), nil
	}
}

// 私钥加密
func (s rsa) PriKeyEncrypt(data, privateKey string) (string, error) {
	privateKey = validPrivate(privateKey)
	return gorsa.PriKeyEncrypt(data, privateKey)
}

// 私钥解密
func (s rsa) PriKeyDecrypt(data, privateKey string) (string, error) {
	privateKey = validPrivate(privateKey)
	if d, err := gorsa.PriKeyDecrypt(data, privateKey); err != nil {
		return "", err
	} else {
		h, _ := hex.DecodeString(d)
		return string(h), nil
	}
}

// 使用RSAWithMD5算法签名
//
//	data: 原文
//	privateKey: 私钥
func (s rsa) SignMd5WithRsa(data, privateKey string) (string, error) {
	privateKey = validPrivate(privateKey)
	return gorsa.SignMd5WithRsa(data, privateKey)
}

// 使用RSAWithMD5验证签名
//
//	data: 原文
//	signData: 签名信息
//	publicKey: 公钥
func (s rsa) VerifySignMd5WithRsa(data, signData, publicKey string) error {
	publicKey = validPublic(publicKey)
	return gorsa.VerifySignMd5WithRsa(data, signData, publicKey)
}

// 使用RSAWithSHA1算法签名
//
//	data: 原文
//	privateKey: 私钥
func (s rsa) SignSha1WithRsa(data, privateKey string) (string, error) {
	privateKey = validPrivate(privateKey)
	return gorsa.SignSha1WithRsa(data, privateKey)
}

// 使用RSAWithSHA1验证签名
//
//	data: 原文
//	signData: 签名信息
//	publicKey: 公钥
func (s rsa) VerifySignSha1WithRsa(data, signData, publicKey string) error {
	publicKey = validPublic(publicKey)
	return gorsa.VerifySignSha1WithRsa(data, signData, publicKey)
}

// 使用RSAWithSHA256算法签名
//
//	data: 原文
//	privateKey: 私钥
func (s rsa) SignSha256WithRsa(data, privateKey string) (string, error) {
	privateKey = validPrivate(privateKey)
	return gorsa.SignSha256WithRsa(data, privateKey)
}

// 使用RSAWithSHA256验证签名
//
//	data: 原文
//	signData: 签名信息
//	publicKey: 公钥
func (s rsa) VerifySignSha256WithRsa(data, signData, publicKey string) error {
	publicKey = validPublic(publicKey)
	return gorsa.VerifySignSha256WithRsa(data, signData, publicKey)
}

// 修正公钥格式
func validPublic(publicKey string) string {
	if !strings.HasPrefix(publicKey, publicPrefix) {
		publicKey = publicPrefix + "\n" + publicKey
	}
	if !strings.HasSuffix(publicKey, publicSuffix) {
		publicKey = publicKey + "\n" + publicSuffix
	}
	return publicKey
}

// 修正私钥格式
func validPrivate(privateKey string) string {
	if !strings.HasPrefix(privateKey, privatePrefix) {
		privateKey = privatePrefix + "\n" + privateKey
	}
	if !strings.HasSuffix(privateKey, privateSuffix) {
		privateKey = privateKey + "\n" + privateSuffix
	}
	return privateKey
}
