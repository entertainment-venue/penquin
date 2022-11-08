package queue

// Queue todo
// queue 接口
// 需要包含哪些
// Queue 队列
type Queue interface {
	Offer([]byte) error
	Poll() ([]byte, error)
	Size() int64
}
