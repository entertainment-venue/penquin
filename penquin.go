package penquin

type PenQuin struct {
	Topics map[string]PQTopic

	PQConfig PQConfig
	PQStats  PQStats
	PQServer PQServer
}

func NewPenquin(pqconfig PQConfig) PenQuin {
	NewPQServer()
}

func (pq *PenQuin) GetTopic(name string) PQTopic {

}

func (pq *PenQuin) GetAllTopic() map[string]PQTopic {

}

func (pq *PenQuin) CreateTopic() error {

}

func (pq *PenQuin) DeleteTopic(name string) error {

}

func (pq *PenQuin) UpdateTopic(name string) error {

}

func (pq *PenQuin) PutMessages(bts []byte) error {

}
