package udputil

import (
	"bufio"
	"net"
	"strconv"
	"time"

	"github.com/zhouweigithub/goutil/iputil"
)

// udp 客户端
type udpClient struct {
	LocalIp       net.IP
	RemoteIp      string
	RemotePort    int
	remoteAddress string //RemoteIp:RemotePort
	conn          net.Conn
	reader        *bufio.Reader
}

// 创建udp客户端
//
//	remoteIp: 远程IP地址
//	remotePort: 远程端口号
func CreateClient(remoteIp string, remotePort int) *udpClient {
	localIp, _ := iputil.GetLocalIp()
	sm := udpClient{
		LocalIp:       localIp,
		RemoteIp:      remoteIp,
		RemotePort:    remotePort,
		remoteAddress: remoteIp + ":" + strconv.Itoa(remotePort),
	}
	return &sm
}

// 发送udp消息
//
//	msg: 消息内容
func (s *udpClient) SendUdp(msg string) error {
	var err error
	if s.conn == nil {
		if s.conn, err = net.Dial("udp", s.remoteAddress); err != nil {
			return err
		}
	}
	if _, err := s.conn.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}

// 接收到消息后的操作
//
// 只能接收到当前连接返回的消息，若消息经过中转，再由其他机器返回，则无法收到。
// 若需要接收任意机器的消息，需要使用server端监听消息
//
//	deal: 接收到消息后的操作
func (s *udpClient) ReceiveUdp(deal func(msg []byte, err error)) {
	go func() {
		data := make([]byte, 4096)
		for {
			if s.conn != nil {
				s.reader = bufio.NewReader(s.conn)
				n, err := s.reader.Read(data)
				deal(data[:n], err)
			} else {
				time.Sleep(time.Second)
			}
		}
	}()
}
