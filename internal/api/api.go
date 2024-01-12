package api

import (
	v1 "github.com/eduardolat/tinylink/internal/api/v1"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/go-chi/chi/v5"
)

func NewRouter(shortener *shortener.Client) chi.Router {
	v1Router := v1.NewRouter(shortener)

	chiRouter := chi.NewRouter()
	chiRouter.Mount("/v1", v1Router)

	return chiRouter
}
