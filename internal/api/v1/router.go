package v1

import (
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func MountRouter(group *echo.Group, shortener *shortener.Shortener) {
	handlers := newHandlers(shortener)

	group.GET("/links", handlers.paginateHandler)
	group.GET("/links/:shortCode", handlers.retrieveHandler)
	group.POST("/links", handlers.shortenHandler)
	group.PUT("/links/:shortCode", handlers.updateHandler)
	group.DELETE("/links/:shortCode", handlers.deleteHandler)
}
