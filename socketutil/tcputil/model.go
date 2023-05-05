package tcputil

type ToServerMsg struct {
	EventType EventType
	ToAddress []string
	Data      interface{}
}

type ToClientMsg struct {
	EventType EventType
	Data      interface{}
}

type EventType int

const (
	SystemMsg EventType = iota + 1
	UserMsg
)
