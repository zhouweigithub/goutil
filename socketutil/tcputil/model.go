package tcputil

// 发送给服务器的消息体
type ToServerMsg struct {
	EventType EventType   // 消息类型
	ToAddress []string    // 目标地址集
	Data      interface{} // 数据
}

// 发送给客户端的消息体
type ToClientMsg struct {
	EventType EventType   // 消息类型
	From      string      // 来源
	Data      interface{} // 数据
}

// 消息类型
type EventType int

const (
	SystemMsg EventType = iota + 1 // 系统消息
	UserMsg                        // 用户消息
)
