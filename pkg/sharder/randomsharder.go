package sharder

import (
	"math/rand"
	"time"
)

var (
	_ Sharder = new(RandomSharder)
)

type RandomSharder struct {
	generator *rand.Rand
}

func NewRandomPartitioner(topic string) Sharder {
	p := new(RandomSharder)
	p.generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return p
}

func (p *RandomSharder) Partition(id string, numPartitions int32) (int32, error) {
	return int32(p.generator.Intn(int(numPartitions))), nil
}
