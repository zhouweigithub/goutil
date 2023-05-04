package udputil

import (
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/zhouweigithub/goutil/iputil"
)

// udp 服务端
type udpServer struct {
	LocalIp      net.IP
	ListenPort   int
	LocalAddress string //LocalIp:ListenPort
	conn         *net.UDPConn
	RemoteAddrs  map[string]*net.UDPAddr
	writeTimeOut time.Duration
}

// 创建udp服务端
//
//	listenPort: 监听的本机端口号
func CreateServer(listenPort int) *udpServer {
	var localIp, _ = iputil.GetLocalIp()
	var s = &udpServer{
		LocalIp:      localIp,
		ListenPort:   listenPort,
		writeTimeOut: time.Second,
		RemoteAddrs:  make(map[string]*net.UDPAddr),
	}
	s.LocalAddress = s.LocalIp.String() + ":" + strconv.Itoa(listenPort)
	return s
}

// 监听本机端口
//
//	deal: 接收到消息后的操作
func (s *udpServer) ListenUdp(deal func(data []byte, remoteAddr *net.UDPAddr, err error)) error {
	var err error
	if s.conn, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   s.LocalIp,
		Port: s.ListenPort,
	}); err != nil {
		return err
	}

	data := make([]byte, 4096)
	go func() {
		for {
			n, remoteAddr, err := s.conn.ReadFromUDP(data)

			var remoteAddrString = remoteAddr.String()
			if _, isOk := s.RemoteAddrs[remoteAddrString]; !isOk {
				s.RemoteAddrs[remoteAddrString] = remoteAddr
			}

			deal(data[:n], remoteAddr, err)
		}
	}()

	return nil
}

// 发送消息
//
//	addrString: 已经建立了连接的地址
//	msg: 消息内容
func (s *udpServer) SendUdp(addrString, msg string) error {
	if s.conn == nil {
		return errors.New("send fail, need listen first")
	}

	if addr, isOk := s.RemoteAddrs[addrString]; isOk {
		s.conn.SetWriteDeadline(time.Now().Add(s.writeTimeOut))
		if _, err := s.conn.WriteToUDP([]byte(msg), addr); err != nil {
			return err
		}
	} else {
		return errors.New("send fail, unconnected with " + addrString)
	}

	return nil
}

// 发送消息到任意地址
//
//	ipString: ip地址
//	port: 端口号
//	msg: 消息内容
func (s *udpServer) SendToAnyUdp(ipString string, port int, msg string) error {
	if s.conn == nil {
		return errors.New("send fail, need listen first")
	}
	var addr = net.UDPAddr{IP: net.ParseIP(ipString), Port: port}
	s.conn.SetWriteDeadline(time.Now().Add(s.writeTimeOut))
	if _, err := s.conn.WriteToUDP([]byte(msg), &addr); err != nil {
		return err
	}
	return nil
}
