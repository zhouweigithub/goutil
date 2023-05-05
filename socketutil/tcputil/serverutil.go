package tcputil

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/zhouweigithub/goutil/iputil"
)

type tcpServer struct {
	LocalIp      net.IP
	ListenPort   int
	LocalAddress string //LocalIp:ListenPort
	RemoteAddrs  map[string]net.Conn
	writeTimeOut time.Duration
}

// 创建udp服务端
//
//	listenPort: 监听的本机端口号
func CreateServer(listenPort int) *tcpServer {
	var localIp, _ = iputil.GetLocalIp()
	var s = &tcpServer{
		LocalIp:      localIp,
		ListenPort:   listenPort,
		writeTimeOut: time.Second,
		RemoteAddrs:  make(map[string]net.Conn),
	}
	s.LocalAddress = s.LocalIp.String() + ":" + strconv.Itoa(listenPort)
	return s
}

func (s *tcpServer) Listen() error {
	listen, err := net.Listen("tcp", s.LocalAddress)
	if err != nil {
		return errors.New("Listen failed, err: " + err.Error())
	}
	fmt.Println("正在监听 " + s.LocalAddress)
	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		fmt.Println("有新的连接：" + conn.RemoteAddr().String())
		go s.process(conn) // 启动一个goroutine来处理客户端的连接请求
	}
}

// TCP Server端测试
// 处理函数
func (s *tcpServer) process(conn net.Conn) {
	defer conn.Close()
	s.RemoteAddrs[conn.RemoteAddr().String()] = conn
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			if errors.Is(err, net.ErrClosed) || errors.Is(err, io.EOF) {
				delete(s.RemoteAddrs, conn.RemoteAddr().String())
				return
			}
			fmt.Println("read from client failed, err: ", err)
			continue
		}
		var msg ToServerMsg
		if err := json.Unmarshal(buf[:n], &msg); err != nil {
			continue
		}

		switch msg.EventType {
		case SystemMsg:
			fmt.Println("收到来自", conn.RemoteAddr().String(), "的消息", msg.Data)
		case UserMsg:
			var toMsg = ToClientMsg{
				Data:      msg.Data,
				EventType: UserMsg,
			}
			bts, _ := json.Marshal(toMsg)
			var toAddress = s.getToConns(msg.ToAddress)
			for _, connItem := range toAddress {
				connItem.SetWriteDeadline(time.Now().Add(time.Second * 5))
				if _, err := connItem.Write(bts); err != nil {
					log.Println(err)
				}
			}
		default:
		}
	}
}

func (s *tcpServer) getToConns(filters []string) []net.Conn {
	var result []net.Conn
	if len(filters) > 0 {
		for _, address := range filters {
			result = append(result, s.RemoteAddrs[address])
		}
	} else {
		for address := range s.RemoteAddrs {
			result = append(result, s.RemoteAddrs[address])
		}
	}
	return result
}
