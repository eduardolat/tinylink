package main

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/tinylink/internal/api"
	"github.com/eduardolat/tinylink/internal/config"
	"github.com/eduardolat/tinylink/internal/datastores/inmemory"
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/shortgens/nanoid"
	"github.com/eduardolat/tinylink/internal/web"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger.Info("🏁 starting TinyLink")
	env := config.GetEnv()

	dataStore := inmemory.NewDataStore()
	err := dataStore.AutoMigrate()
	if err != nil {
		logger.FatalError(
			"failed to auto migrate data store",
			"error",
			err,
		)
	}

	shortGen := nanoid.NewShortGen()
	shortenerClient := shortener.NewShortener(dataStore, shortGen)

	appRouter := chi.NewRouter()
	apiRouter := api.NewRouter(shortenerClient)
	webRouter := web.NewRouter(shortenerClient)

	appRouter.Mount("/api", apiRouter)
	appRouter.Mount("/", webRouter)

	port := fmt.Sprintf(":%d", env.PORT)
	logger.Info("🚀 starting HTTP server", "port", port)
	err = http.ListenAndServe(port, appRouter)
	if err != nil {
		logger.FatalError(
			"failed to start HTTP server",
			"error",
			err,
		)
	}
}
