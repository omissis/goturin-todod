package uuid

import (
	"github.com/google/uuid"
)

func NewGoogleGenerator() *GoogleGenerator {
	return &GoogleGenerator{}
}

type GoogleGenerator struct{}

func (g *GoogleGenerator) Generate() string {
	return uuid.New().String()
}
