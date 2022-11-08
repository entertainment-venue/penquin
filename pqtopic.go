package penquin

import (
	"github.com/entertainment-venue/penquin/sharder"
)

type PQTopic struct {
	Name      string
	ShardsMap map[string]*PQShard
	Info      TopicInfo
	sd        sharder.Sharder

	messageCount int64
	messageBytes int64
}

type TopicInfo struct {
	Name       string `json:"name"`
	ShardCnt   int    `json:"shard_cnt"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func NewPQTopic() (*PQTopic, error) {

}

func (t *PQTopic) AddMessage(score float64, content []byte) error {

}
