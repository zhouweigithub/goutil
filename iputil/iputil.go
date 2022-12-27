package iputil

import (
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/zhouweigithub/goutil/errutil"
	"github.com/zhouweigithub/goutil/logutil"
	"github.com/zhouweigithub/goutil/randutil"
	"github.com/zhouweigithub/goutil/webutil"
)

var ipServerList = [...]string{"https://ipv4.ddnspod.com", "https://ipecho.net/plain", "https://ipinfo.io/ip"}

// 获取本地计算机的远程IP地址
func GetRemoteIp(timeout time.Duration) string {
	defer errutil.CatchError()
	var randIndex = randutil.GetRandInt(0, len(ipServerList)-1)
	html, _, err := webutil.GetWithTimeOut(ipServerList[randIndex], timeout)
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

// 获取本地IP地址 利用udp
func GetLocalIp() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return nil, err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
