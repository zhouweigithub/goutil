package iputil

import (
	"regexp"
	"strings"
	"time"

	errutil "github.com/zhouweigithub/goutil/errUtil"
	logutil "github.com/zhouweigithub/goutil/logUtil"
	randutil "github.com/zhouweigithub/goutil/randUtil"
	webutil "github.com/zhouweigithub/goutil/webUtil"
)

var ipServerList = [...]string{"https://ipv4.ddnspod.com", "https://ipecho.net/plain", "https://ipinfo.io/ip"}

//获取本地计算机的远程IP地址
func GetRemoteIp(timeout time.Duration) string {
	defer errutil.CatchError()
	var randIndex = randutil.GetRandInt(0, len(ipServerList)-1)
	html, err, _ := webutil.GetWithTimeOut(ipServerList[randIndex], timeout)
	if err != nil {
		logutil.Error("获取本机远程IP失败：" + err.Error())
		return ""
	} else {
		var ip = strings.Trim(strings.Trim(html, "\n"), " ")
		if IsIp(ip) {
			return ip
		} else {
			return ""
		}
	}
}

// 字符串是否为合法IP地址
func IsIp(ipstring string) bool {
	var isOk, _ = regexp.Match(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`, []byte(ipstring))
	return isOk
}
