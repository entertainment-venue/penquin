package orderqueue

import "github.com/go-redis/redis"

var (
	_ OrderQueue = new(RedisOrderQueue)
)

type RedisOrderQueue struct {
	redis.Client
}

func (r RedisOrderQueue) Push(name string, msg Message) error {
	r.ZAdd()
}

func (r RedisOrderQueue) MultiPush(name string, msgs []Message) error {
	r.ZAdd()
}

func (r RedisOrderQueue) PushNX(name string, msg Message) error {
	panic("implement me")
}

func (r RedisOrderQueue) MultiPushNX(name string, msgs []Message) error {
	panic("implement me")
}

func (r RedisOrderQueue) Pop(name string, count int) ([]Message, error) {
	panic("implement me")
}

func (r RedisOrderQueue) Del(name string, msg Message) error {
	panic("implement me")
}

func (r RedisOrderQueue) MulDel(name string, msgs []Message) error {
	panic("implement me")
}

func (r RedisOrderQueue) Size(name string) int {
	panic("implement me")
}

func (r RedisOrderQueue) Ack(name, id string) error {
	panic("implement me")
}

func (r RedisOrderQueue) Flush(name string) error {
	panic("implement me")
}
