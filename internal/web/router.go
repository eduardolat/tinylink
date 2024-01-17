package web

import (
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func MountRouter(group *echo.Group, shortener *shortener.Shortener) {
	handlers := newHandlers(shortener)

	group.GET("/:shortCode", handlers.redirect)
}
