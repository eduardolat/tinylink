package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

/*
	More info of the algorithm:

	https://github.com/ai/nanoid
	https://github.com/matoous/go-nanoid
*/

/*
	With the default alphabet and length, 2M IDs are needed in order
	to have a 1% probability of at least one collision.

	Calculator: https://zelark.github.io/nano-id-cc/

	The shortener will check if the generated ID already exists in the
	database. If it does, it will generate a new one until it finds one
	that doesn't exist in the database.

	Examples with the default alphabet and length:
	- mjPOtEIV
	- N2Yc3ZAo
	- 5JlVWlxh
	- 2ZEvjRzt
*/

const (
	// DefaultLength is the default length of the generated NanoID
	DefaultLength = 8
	// DefaultAlphabet is the default alphabet that will be used to generate
	// the NanoID
	DefaultAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// NanoIDShortCodeGenerator is a short code generator that uses NANOIDs
// to generate the unique short codes
type NanoIDShortCodeGenerator struct {
	alphabet string
	length   int
}

// NewNanoIDShortCodeGeneratorOptions is a struct that contains the options
// for the NewNANOIDShortCodeGenerator function
type NewNanoIDShortCodeGeneratorOptions struct {
	// Alphabet is a custom alphabet that will be used to generate
	// the NanoID, the default value is used if the value is empty
	Alphabet string

	// Length is the length of the generated NanoID, the default value
	// is used if the value is 0
	Length int
}

// NewNanoIDShortCodeGenerator returns a new NANOIDShortCodeGenerator
func NewNanoIDShortCodeGenerator(
	options NewNanoIDShortCodeGeneratorOptions,
) *NanoIDShortCodeGenerator {
	cg := &NanoIDShortCodeGenerator{
		alphabet: DefaultAlphabet,
		length:   DefaultLength,
	}

	if options.Alphabet != "" {
		cg.alphabet = options.Alphabet
	}

	if options.Length != 0 {
		cg.length = options.Length
	}

	return cg
}

// Generate generates a new unique short code for the URL shortener
func (g *NanoIDShortCodeGenerator) Generate() (string, error) {
	u, err := gonanoid.Generate(g.alphabet, g.length)
	if err != nil {
		return "", err
	}
	return u, nil
}
