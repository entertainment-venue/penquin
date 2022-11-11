package penquin

import (
	"encoding/json"
)

type PQMessage struct {
	UUID      string  `json:"UUID"`
	Score     float64 `json:"Score"`
	Content   string  `json:"Content"`
	CreatedAt string  `json:"CreatedAt"`
	UpdatedAt string  `json:"UpdatedAt"`
}

func (m *PQMessage) ToJsonBytes() []byte {
	bts, _ := json.Marshal(m)
	return bts
}

func GeneratePQMessageFromJsonStr(str string) (*PQMessage, error) {
	msg := &PQMessage{}
	if err := json.Unmarshal([]byte(str), &msg); err != nil {
		return nil, err
	}
	return msg, nil
}
