package penquin

import (
	"github.com/entertainment-venue/penquin/pkg/kvstore"
	queue2 "github.com/entertainment-venue/penquin/pkg/queue"
	"time"

	uuid "github.com/satori/go.uuid"
)

type PQShard struct {
	readyQueue                queue2.Queue
	orderQueue                queue2.OrderedQueue
	runtimeStore, backUpStore kvstore.KVStore

	close chan struct{}
	//duration
}

func NewPQShard() *PQShard {
	rstore, _ := kvstore.NewRedisKVStore()
	return &PQShard{
		runtimeStore: rstore,
		backUpStore:  nil,
	}
}

func (s *PQShard) AddMessage(score float64, content []byte) error {
	msgId := uuid.NewV4().String()
	msg := PQMessage{
		UUID:      msgId,
		Score:     score,
		Content:   string(content),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	if err := s.runtimeStore.Put([]byte(msgId), msg.ToJsonBytes()); err != nil {
		return err
	}
	return s.orderQueue.OfferWithScore(content, score)
}

func (s *PQShard) pushToReadyQ() {
LOOP:
	for {
		select {
		case <-s.close:
			return
		default:
			timeNowUnix := float64(time.Now().Unix())
			msgs, err := s.orderQueue.PollWithScore(0, timeNowUnix)
			if err != nil {
				SError("poll messages from shard failed, error is %+v", err)
				continue
			}
			for _, v := range msgs {
				value, err := s.runtimeStore.Get(v)
				if err != nil {
					SError("push message to ready queue failed, message is %s, error is %+v", string(v), err)
					goto LOOP
				}
				if err := s.readyQueue.Offer(value); err != nil {
					SError("push message to ready queue failed, message is %s, error is %+v", string(v), err)
					goto LOOP
				}
				if err := s.backUpStore.Put(v, value); err != nil {
					SError("push message to ready queue failed, message is %s, error is %+v", string(v), err)
					goto LOOP
				}
				flag, err := s.runtimeStore.Del(v)
				if err != nil {
					SError("push message to ready queue failed, message is %s, delete status is %t, error is %+v", string(v), flag, err)
					goto LOOP
				}
			}
			cnt, err := s.orderQueue.RemoveWithScore(0, timeNowUnix)
			if err != nil {
				SError("push message to ready queue failed, message is %s, delete status is %t, error is %+v", err)
				continue
			}
			if cnt != int64(len(msgs)) {
				SError("push message to ready queue failed, message is %s, delete status is %t, error is %+v", err)
			}
		}
	}
}
