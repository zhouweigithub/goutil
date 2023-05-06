package tcputil

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/zhouweigithub/goutil/iputil"
)

// TCP服务端对象
type tcpServer struct {
	LocalIp      net.IP              // 本地IP地址
	ListenPort   int                 // 监听端口号
	LocalAddress string              // LocalIp:ListenPort
	RemoteAddrs  map[string]net.Conn // 所有连接集合（ip地址/conn）
}

// 创建udp服务端
//
//	listenPort: 监听的本机端口号
func CreateServer(listenPort int) *tcpServer {
	var localIp, _ = iputil.GetLocalIp()
	var s = &tcpServer{
		LocalIp:     localIp,
		ListenPort:  listenPort,
		RemoteAddrs: make(map[string]net.Conn),
	}
	s.LocalAddress = s.LocalIp.String() + ":" + strconv.Itoa(listenPort)
	return s
}

// 监听本地计算机，等待用户连接
func (s *tcpServer) Listen() error {
	listen, err := net.Listen("tcp", s.LocalAddress)
	if err != nil {
		return errors.New("Listen failed, err: " + err.Error())
	}
	fmt.Println("服务启动成功")
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

// 处理当前连接的收发消息
func (s *tcpServer) process(conn net.Conn) {
	defer conn.Close()

	s.notifyUserOnLine(conn.RemoteAddr().String())
	s.notifyOnLineUser(conn)

	s.RemoteAddrs[conn.RemoteAddr().String()] = conn
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			// if err == net.ErrClosed || err == io.EOF {
			delete(s.RemoteAddrs, conn.RemoteAddr().String())
			fmt.Println(conn.RemoteAddr().String() + " 已断开连接")
			return
			// }
		}
		var msg ToServerMsg
		if err := json.Unmarshal(buf[:n], &msg); err != nil {
			continue
		}

		switch msg.EventType {
		case SystemMsg:
			fmt.Println("收到来自", conn.RemoteAddr().String(), "的系统消息", msg.Data)
		case UserMsg:
			fmt.Println("收到来自", conn.RemoteAddr().String(), "的用户消息", msg.Data)
			var toMsg = ToClientMsg{
				Data:      msg.Data,
				EventType: UserMsg,
				From:      conn.RemoteAddr().String(),
			}
			bts, _ := json.Marshal(toMsg)
			var toAddress = s.getToConns(msg.ToAddress, conn.RemoteAddr().String())
			for _, connItem := range toAddress {
				if err := connItem.SetWriteDeadline(time.Now().Add(time.Second * 5)); err != nil {
					connItem.Close()
					log.Println(err)
				}
				if _, err := connItem.Write(bts); err != nil {
					connItem.Close()
					log.Println(err)
				}
			}
		default:
		}
	}
}

// 获取需要发送消息的用户连接
//
//	filters: 过滤的ip信息
//	currentAddr: 当前用户的ip信息
func (s *tcpServer) getToConns(filters []string, currentAddr string) []net.Conn {
	var result []net.Conn
	if len(filters) > 0 {
		for _, address := range filters {
			if conn, isok := s.RemoteAddrs[address]; isok {
				result = append(result, conn)
			}
		}
	} else {
		for address := range s.RemoteAddrs {
			if address == currentAddr {
				continue
			}
			result = append(result, s.RemoteAddrs[address])
		}
	}
	return result
}

// 获取所有在线用户集
func (s *tcpServer) getOnlineUsers() []string {
	var result []string
	for address := range s.RemoteAddrs {
		result = append(result, address)
	}
	return result
}

// 通知其他用户新用户上线
func (s *tcpServer) notifyUserOnLine(newAddr string) {
	var toMsg = ToClientMsg{
		Data:      newAddr + "上线了",
		EventType: SystemMsg,
	}
	bts, _ := json.Marshal(toMsg)
	for addr, connItem := range s.RemoteAddrs {
		if addr == newAddr {
			continue
		}
		if _, err := connItem.Write(bts); err != nil {
			connItem.Close()
			log.Println(err)
		}
	}
}

// 通知用户在线情况
func (s *tcpServer) notifyOnLineUser(conn net.Conn) {
	var toMsg = ToClientMsg{
		Data:      "当前在线用户(" + strconv.Itoa(len(s.RemoteAddrs)+1) + ")" + strings.Join(s.getOnlineUsers(), "\n"),
		EventType: SystemMsg,
	}
	bts, _ := json.Marshal(toMsg)
	if _, err := conn.Write(bts); err != nil {
		conn.Close()
		log.Println(err)
	}
}
