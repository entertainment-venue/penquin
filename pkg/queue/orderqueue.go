package queue

// OrderedQueue 有序队列
type OrderedQueue interface {
	OfferWithScore(value []byte, score float64) error
	PollWithScore(start float64, end float64) ([][]byte, error)
	RemoveWithScore(start float64, end float64) (int64, error)
	PollPeak() ([]byte, error)
	PollTail() ([]byte, error)
	Size() int64
}
