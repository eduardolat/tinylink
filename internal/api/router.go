package api

import (
	v1 "github.com/eduardolat/tinylink/internal/api/v1"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func MountRouter(group *echo.Group, shortener *shortener.Shortener) {
	v1Group := group.Group("/v1")
	v1.MountRouter(v1Group, shortener)
}
