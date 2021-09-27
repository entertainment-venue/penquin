package orderqueue

type OrderQueue interface {
	Push(name string, msg Message) error
	MultiPush(name string, msgs []Message) error
	PushNX(name string, msg Message) error
	MultiPushNX(name string, msgs []Message) error
	Pop(name string, count int) ([]Message, error)
	Del(name string, msg Message) error
	MulDel(name string, msgs []Message) error
	Size(name string) int
	Ack(name, id string) error
	Flush(name string) error
}

type Message struct {
	Id         string  `json:"id"`
	Topic      string  `json:"topic"`
	Delay      float64 `json:"delay"`
	Expire     float64 `json:"expire"`
	CreateTime float64 `json:"create_time"`
	Content    string  `json:"content"`
}

