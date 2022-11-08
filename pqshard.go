package penquin

import (
	"github.com/entertainment-venue/penquin/kvstore"
	uuid "github.com/satori/go.uuid"
	"time"

	"github.com/entertainment-venue/penquin/queue"
	"go.uber.org/zap"
)

type PQShard struct {
	readyQueue                queue.Queue
	orderQueue                queue.OrderedQueue
	runtimeStore, backUpStore kvstore.KVStore
	t                         time.Timer
	close                     chan struct{}
	lg                        zap.SugaredLogger
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
	for {
		msgs, err := s.orderQueue.PollWithScore(0, float64(time.Now().Unix()))
		if err != nil {
			s.lg.Errorf("poll messages from shard failed, error is %+v", err)
			continue
		}
		for _, v := range msgs {
			value, err := s.runtimeStore.Get(v)
			if err != nil {
				s.lg.Errorf("push message to ready queue failed, message is %s, error is %+v", string(v), err)
				continue
			}
			if err := s.readyQueue.Offer(value); err != nil {
				s.lg.Errorf("push message to ready queue failed, message is %s, error is %+v", string(v), err)
				continue
			}
			if err := s.backUpStore.Put(v, value); err != nil {
				
			}
			flag, err := s.runtimeStore.Del(v)
			if err != nil {
				s.lg.Errorf("push message to ready queue failed, message is %s, delete status is %t, error is %+v", string(v), flag, err)
				continue
			}
			flag
		}
	}
}
