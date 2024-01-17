package web

import (
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func MountRouter(group *echo.Group, shortener *shortener.Shortener) {
	handlers := newHandlers(shortener)

	group.Any("/", handlers.notFoundHandler)
	group.GET("/:shortCode", handlers.redirectHandler)
}
