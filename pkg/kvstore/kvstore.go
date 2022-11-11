package kvstore

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ KVStore = new(rediskv)

type KVStore interface {
	Put([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) (bool, error)
}

func NewRedisKVStore() (*rediskv, error) {
	return nil, nil
}

type rediskv struct {
	rdb *redis.Client

	expireDuration time.Duration
}

func (r *rediskv) Put(key []byte, value []byte) error {
	return r.rdb.SetNX(context.Background(), string(key), string(value), 0).Err()
}

func (r *rediskv) Get(key []byte) ([]byte, error) {
	value, err := r.rdb.Get(context.Background(), string(key)).Result()
	if err != nil {
		return nil, err
	}
	return []byte(value), nil
}

func (r *rediskv) Del(key []byte) (bool, error) {
	cnt, err := r.rdb.Del(context.Background(), string(key)).Result()
	if err != nil {
		return false, err
	}
	return cnt == 1, nil
}
