package uuid

import (
	googleuuid "github.com/google/uuid"
)

// UUIDShortCodeGenerator is a short code generator that uses UUIDs
// to generate the unique short codes
type UUIDShortCodeGenerator struct{}

// NewUUIDShortCodeGenerator returns a new UUIDShortCodeGenerator
func NewUUIDShortCodeGenerator() *UUIDShortCodeGenerator {
	return &UUIDShortCodeGenerator{}
}

// Generate generates a new unique short code for the URL shortener
func (g *UUIDShortCodeGenerator) Generate() (string, error) {
	u, err := googleuuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
