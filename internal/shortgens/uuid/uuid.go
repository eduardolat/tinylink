package uuid

import (
	"strings"

	googleuuid "github.com/google/uuid"
)

// ShortGen is a short code generator that uses UUIDs
// to generate the unique short codes
type ShortGen struct {
	removeDashes bool
}

// ShortGenOpts is a struct that contains the options
// for the short code generator
type ShortGenOpts struct {
	// RemoveDashes is a flag that indicates whether the dashes should be removed
	// from the generated UUIDs
	RemoveDashes bool
}

// NewShortGen returns a new UUID ShortGen
// Options are optional and the default values will be used if they are not provided.
func NewShortGen(
	options ...ShortGenOpts,
) *ShortGen {
	pickedOptions := ShortGenOpts{
		RemoveDashes: false,
	}
	if len(options) > 0 {
		pickedOptions = options[0]
	}

	return &ShortGen{
		removeDashes: pickedOptions.RemoveDashes,
	}
}

// Generate generates a new unique short code for the URL shortener
func (g *ShortGen) Generate() (string, error) {
	u, err := googleuuid.NewRandom()
	if err != nil {
		return "", err
	}

	shortCode := u.String()
	if g.removeDashes {
		shortCode = strings.ReplaceAll(shortCode, "-", "")
	}

	return shortCode, nil
}
