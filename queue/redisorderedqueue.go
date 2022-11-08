package queue

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	_ OrderedQueue = new(redisorderedqueue)

	WrongTypeErr = errors.New("wrong type")
)

type redisorderedqueue struct {
	queueName string

	rdb *redis.Client
}

func (r *redisorderedqueue) PollWithScore(start float64, end float64) ([][]byte, error) {
	zres, err := r.rdb.ZRangeByScore(context.Background(), r.queueName, &redis.ZRangeBy{
		Min:    strconv.FormatFloat(start, 'f', -1, 64),
		Max:    strconv.FormatFloat(end, 'f', -1, 64),
		Offset: 0,
		Count:  int64(1),
	}).Result()
	if err != nil {
		return nil, err
	}
	if len(zres) <= 0 {
		return nil, nil
	}
	var msgs [][]byte
	for _, res := range zres {
		msg, err := GenerateFromJsonStr(res)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, []byte(msg.Content))
	}
	return msgs, nil
}

func (r *redisorderedqueue) RemoveWithScore(start float64, end float64) (int64, error) {
	zremcnt, err := r.rdb.ZRemRangeByScore(context.Background(), r.queueName,
		strconv.FormatFloat(start, 'f', -1, 64),
		strconv.FormatFloat(end, 'f', -1, 64), ).Result()
	if err != nil {
		return 0, err
	}
	return zremcnt, nil
}

func (r *redisorderedqueue) OfferWithScore(bytes []byte, f float64) error {
	message := &RedisOrderedMessage{
		RedisMessage: RedisMessage{
			UUID:      uuid.NewV4().String(),
			Content:   string(bytes),
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		},
		Score: f,
	}
	return r.rdb.ZAdd(context.Background(), r.queueName, &redis.Z{
		Score:  f,
		Member: message.ToJsonStr(),
	}).Err()
}

func (r *redisorderedqueue) PollPeak() ([]byte, error) {
	results, err := r.rdb.ZPopMax(context.Background(), r.queueName).Result()
	if err != nil {
		return nil, err
	}
	if len(results) <= 0 {
		return []byte{}, nil
	}
	if msgStr, ok := results[0].Member.(string); ok {
		msg, err := GenerateFromJsonStr(msgStr)
		if err != nil {
			return nil, err
		}
		return []byte(msg.Content), nil
	}
	return nil, WrongTypeErr
}

func (r *redisorderedqueue) PollTail() ([]byte, error) {
	results, err := r.rdb.ZPopMin(context.Background(), r.queueName).Result()
	if err != nil {
		return nil, err
	}
	if len(results) <= 0 {
		return []byte{}, nil
	}
	if msgStr, ok := results[0].Member.(string); ok {
		msg, err := GenerateFromJsonStr(msgStr)
		if err != nil {
			return nil, err
		}
		return []byte(msg.Content), nil
	}
	return nil, WrongTypeErr
}

func (r *redisorderedqueue) Size() int64 {
	return r.rdb.ZCard(context.Background(), r.queueName).Val()
}

type RedisOrderedMessage struct {
	RedisMessage

	Score float64 `json:"Score"`
}

func (m *RedisOrderedMessage) ToJsonStr() string {
	bts, _ := json.Marshal(m)
	return string(bts)
}

func GenerateFromJsonStr(str string) (*RedisOrderedMessage, error) {
	msg := &RedisOrderedMessage{}
	if err := json.Unmarshal([]byte(str), &msg); err != nil {
		return nil, err
	}
	return msg, nil
}
