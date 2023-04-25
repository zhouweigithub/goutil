package iputil

import (
	"net"
	"regexp"
	"strings"

	"github.com/zhouweigithub/goutil/errutil"
	"github.com/zhouweigithub/goutil/logutil"
	"github.com/zhouweigithub/goutil/randutil"
	"github.com/zhouweigithub/goutil/webutil"
)

// 获取本机公网IP的域名集
var ipServerList = [...]string{"https://ipv4.ddnspod.com", "https://ipecho.net/plain", "https://ipinfo.io/ip"}

// 获取本地计算机的公网IP地址
//
//	timeout: 超时时间（秒）
func GetRemoteIp(timeout int) string {
	defer errutil.CatchError()
	var randIndex = randutil.GetRandInt(0, len(ipServerList)-1)
	html, _, _, err := webutil.GetWeb(ipServerList[randIndex], nil, nil, "", timeout)
	if err != nil {
		logutil.Error("获取本机公网IP失败：" + err.Error())
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
