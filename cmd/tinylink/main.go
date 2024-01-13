package main

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/api"
	"github.com/eduardolat/tinylink/internal/datastores/inmemory"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/shortgens/nanoid"
	"github.com/eduardolat/tinylink/internal/web"
	"github.com/go-chi/chi/v5"
)

func main() {
	shortGen := nanoid.NewShortGen()
	dataStore := inmemory.NewInMemoryDataStore()
	shortenerClient := shortener.NewClient(shortGen, dataStore)

	appRouter := chi.NewRouter()
	apiRouter := api.NewRouter(shortenerClient)
	webRouter := web.NewRouter(shortenerClient)

	appRouter.Mount("/api", apiRouter)
	appRouter.Mount("/", webRouter)

	http.ListenAndServe(":8080", appRouter)
}
