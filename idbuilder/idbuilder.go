package idbuilder

type IdBuilder interface {
	Generate(id string) string
}

