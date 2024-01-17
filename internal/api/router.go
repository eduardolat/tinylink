package api

import (
	v1 "github.com/eduardolat/tinylink/internal/api/v1"
	"github.com/eduardolat/tinylink/internal/middleware"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func MountRouter(
	group *echo.Group,
	mid *middleware.Middleware,
	shortener *shortener.Shortener,
) {
	group.Use(mid.BasicAuthIfEnabled)

	v1Group := group.Group("/v1")
	v1.MountRouter(v1Group, shortener)
}
