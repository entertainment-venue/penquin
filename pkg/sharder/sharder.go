package sharder

type Sharder interface {
	Partition(id string, numPartitions int32) (int32, error)
}
