package penquin

type PenQuin struct {
	Topics map[string]PQTopic

	PQConfig PQConfig
	PQStats  PQStats
	PQServer PQServer
}

func NewPenquin(pqconfig PQConfig) *PenQuin {
	//NewPQServer()
	return nil
}

func (pq *PenQuin) GetTopic(name string) *PQTopic {
	return nil
}

func (pq *PenQuin) GetAllTopic() map[string]PQTopic {
	return nil
}

func (pq *PenQuin) CreateTopic() error {
	return nil
}

func (pq *PenQuin) DeleteTopic(name string) error {
	return nil
}

func (pq *PenQuin) UpdateTopic(name string) error {
	return nil
}

func (pq *PenQuin) PutMessages(bts []byte) error {
	return nil
}
