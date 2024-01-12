package v1

import (
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/go-chi/chi/v5"
)

type router struct {
	shortener *shortener.Client
}

func NewRouter(shortener *shortener.Client) chi.Router {
	r := &router{
		shortener: shortener,
	}

	chiRouter := chi.NewRouter()
	chiRouter.HandleFunc("/shorten", r.shorten)
	chiRouter.HandleFunc("/retrieve", r.retrieve)

	return chiRouter
}
