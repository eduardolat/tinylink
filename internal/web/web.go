package web

import (
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/go-chi/chi/v5"
)

type ruoter struct {
	shortener *shortener.Client
}

func NewRouter(shortener *shortener.Client) chi.Router {
	r := &ruoter{
		shortener: shortener,
	}

	chiRouter := chi.NewRouter()
	chiRouter.HandleFunc("/{shortCode}", r.redirect)

	return chiRouter
}
