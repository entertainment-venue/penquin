package queue

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var _ Queue = new(redisqueue)

type redisqueue struct {
	queueName string

	rdb *redis.Client
}

func (r *redisqueue) Offer(bytes []byte) error {
	return r.rdb.LPush(context.Background(), r.queueName, string(bytes)).Err()
}

func (r *redisqueue) Poll() ([]byte, error) {
	msg, err := r.rdb.LPop(context.Background(), r.queueName).Result()
	if err != nil {
		return nil, err
	}
	return []byte(msg), nil
}

func (r *redisqueue) Size() int64 {
	return r.rdb.LLen(context.Background(), r.queueName).Val()
}

type RedisMessage struct {
	UUID      string `json:"UUID"`
	Content   string `json:"Content"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}
