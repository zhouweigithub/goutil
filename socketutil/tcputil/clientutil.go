package tcputil

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/zhouweigithub/goutil/iputil"
)

type tcpClient struct {
	LocalIp       net.IP
	RemoteIp      string
	RemotePort    int
	remoteAddress string //RemoteIp:RemotePort
	conn          net.Conn
	OnReceivedMsg func(*ToClientMsg)
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

func (t *tcpClient) Connect() error {
	conn, err := net.Dial("tcp", t.remoteAddress)
	if err != nil {
		return err
	}
	t.conn = conn
	t.ReceivedMsg()
	return nil
}

func (t *tcpClient) Close() error {
	return t.conn.Close()
}

func (t *tcpClient) ReceivedMsg() {
	buf := [1024]byte{}
	go func() {
		for {
			n, err := t.conn.Read(buf[:])
			if err != nil {
				fmt.Println("recv failed, err:", err)
				if errors.Is(err, net.ErrClosed) {
					if t.Connect() != nil {
						time.Sleep(time.Second * 10)
					}
				}
				return
			}
			var msg ToClientMsg
			if err := json.Unmarshal(buf[:n], &msg); err != nil {
				fmt.Println(err)
			} else {
				if t.OnReceivedMsg != nil {
					t.OnReceivedMsg(&msg)
				}
			}
		}
	}()
}

func (t *tcpClient) SendSysMsg(msg string) {
	var data = ToServerMsg{
		EventType: SystemMsg,
		Data:      msg,
	}
	t.send(data)
}

func (t *tcpClient) SendUserMsg(msg string, tos []string) {
	var data = ToServerMsg{
		EventType: UserMsg,
		Data:      msg,
		ToAddress: tos,
	}
	t.send(data)
}

func (t *tcpClient) send(data interface{}) {
	bts, _ := json.Marshal(data)
	t.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	if _, err := t.conn.Write(bts); err != nil {
		log.Println(err)
	}
}
