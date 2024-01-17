package v1

import "github.com/eduardolat/tinylink/internal/shortener"

type handlers struct {
	shortener *shortener.Shortener
}

func newHandlers(shortener *shortener.Shortener) *handlers {
	return &handlers{
		shortener: shortener,
	}
}
