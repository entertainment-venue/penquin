package idbuilder

import "github.com/google/uuid"

type uuidGenerator struct {
}

func (u uuidGenerator) Generate() string {
	return uuid.New().String()
}
