package main

import (
	"fmt"

	"github.com/eduardolat/tinylink/internal/api"
	"github.com/eduardolat/tinylink/internal/config"
	"github.com/eduardolat/tinylink/internal/datastores/inmemory"
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/eduardolat/tinylink/internal/middleware"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/shortgens/nanoid"
	"github.com/eduardolat/tinylink/internal/web"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.Info("✂️  starting TinyLink")
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
	mid := middleware.NewMiddleware(env)

	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	webGroup := app.Group("")
	web.MountRouter(webGroup, mid, shortenerClient)

	apiGroup := app.Group("/api")
	api.MountRouter(apiGroup, mid, shortenerClient)

	port := fmt.Sprintf(":%d", *env.PORT)
	logger.Info("🚀 HTTP server started", "port", port)

	err = app.Start(port)
	if err != nil {
		logger.FatalError(
			"failed to start HTTP server",
			"error",
			err,
		)
	}
}
