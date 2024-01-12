package shortener

// ShortCodeGenerator is the interface that wraps the Generate method
// that will be implemented by all the short code generators
type ShortCodeGenerator interface {
	// Generate generates a new unique short code for the URL shortener
	Generate() (string, error)
}
