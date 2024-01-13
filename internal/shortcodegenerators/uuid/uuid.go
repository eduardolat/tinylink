package uuid

import (
	googleuuid "github.com/google/uuid"
)

// UUIDShortCodeGenerator is a short code generator that uses UUIDs
// to generate the unique short codes
type UUIDShortCodeGenerator struct {
	removeDashes bool
}

// NewUUIDShortCodeGeneratorOptions is a struct that contains the options
// for the NewUUIDShortCodeGenerator function
type NewUUIDShortCodeGeneratorOptions struct {
	// RemoveDashes is a flag that indicates whether the dashes should be removed
	// from the generated UUIDs
	RemoveDashes bool
}

// NewUUIDShortCodeGenerator returns a new UUIDShortCodeGenerator
func NewUUIDShortCodeGenerator(
	options NewUUIDShortCodeGeneratorOptions,
) *UUIDShortCodeGenerator {
	return &UUIDShortCodeGenerator{
		removeDashes: options.RemoveDashes,
	}
}

// Generate generates a new unique short code for the URL shortener
func (g *UUIDShortCodeGenerator) Generate() (string, error) {
	u, err := googleuuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
