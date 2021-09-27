package penquin

import (
	"github.com/entertainment-venue/penquin/orderqueue"
	"github.com/entertainment-venue/penquin/sharder"
)

type PQTopic struct {
	Name      string
	ShardsMap map[string]orderqueue.OrderQueue
	Info      TopicInfo
	sd   sharder.Sharder

	messageCount int64
	messageBytes int64
}

type TopicInfo struct {
	Name       string `json:"name"`
	Shards     int    `json:"shards"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
