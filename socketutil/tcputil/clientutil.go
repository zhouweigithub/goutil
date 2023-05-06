package tcputil

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/zhouweigithub/goutil/iputil"
	"github.com/zhouweigithub/goutil/sliceutil"
)

// TCP客户端对象
type tcpClient struct {
	LocalIp       net.IP             // 本地IP地址
	RemoteIp      string             // 服务器IP地址
	RemotePort    int                // 服务器端口号
	remoteAddress string             // RemoteIp:RemotePort
	conn          net.Conn           // 与服务器之间的连接
	OnReceivedMsg func(*ToClientMsg) // 接收消息后的处理函数
}

// 创建tcp客户端
//
//	remoteIp: 远程IP地址
//	remotePort: 远程端口号
func CreateClient(remoteIp string, remotePort int) *tcpClient {
	localIp, _ := iputil.GetLocalIp()
	sm := tcpClient{
		LocalIp:       localIp,
		RemoteIp:      remoteIp,
		RemotePort:    remotePort,
		remoteAddress: remoteIp + ":" + strconv.Itoa(remotePort),
	}
	return &sm
}

// 连接服务器并接收消息
func (t *tcpClient) Connect() error {
	if err := t.connect(); err != nil {
		return err
	}
	t.receiveMsg()
	return nil
}

// 连接服务器
func (t *tcpClient) connect() error {
	conn, err := net.Dial("tcp", t.remoteAddress)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}

// 关闭连接
func (t *tcpClient) Close() error {
	return t.conn.Close()
}

// 接收消息
func (t *tcpClient) receiveMsg() {
	buf := [1024]byte{}
	go func() {
		defer func() {
			fmt.Println("连接已断开")
		}()
		for {
		start:
			n, err := t.conn.Read(buf[:])
			if err != nil {
				fmt.Println("recv failed, err:", err)
				for i := 0; i < 3; i++ {
					fmt.Println("尝试第" + strconv.Itoa(i+1) + "次重新连接服务器...")
					if t.connect() != nil {
						time.Sleep(time.Second * 10)
					} else {
						fmt.Println("服务器重新连接成功")
						goto start
					}
				}
				return
			} else {
				var msg ToClientMsg
				if err := json.Unmarshal(buf[:n], &msg); err != nil {
					fmt.Println(err)
				} else {
					if t.OnReceivedMsg != nil {
						t.OnReceivedMsg(&msg)
					}
				}
			}
		}
	}()
}

// 给系统发送消息
func (t *tcpClient) SendSysMsg(msg string) {
	msg = strings.TrimSpace(msg)
	var data = ToServerMsg{
		EventType: SystemMsg,
		Data:      msg,
	}
	t.send(data)
}

// 给用户发送消息
func (t *tcpClient) SendUserMsg(msg string, tos ...string) {
	msg = strings.TrimSpace(msg)
	// 移除空数据
	tos = sliceutil.Remove(tos, func(item *string) bool { return *item == "" })
	var data = ToServerMsg{
		EventType: UserMsg,
		Data:      msg,
		ToAddress: tos,
	}
	t.send(data)
}

// 发送消息
func (t *tcpClient) send(data interface{}) {
	bts, _ := json.Marshal(data)
	t.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	if _, err := t.conn.Write(bts); err != nil {
		log.Println(err)
	}
}
