package v1

import (
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func MountRouter(group *echo.Group, shortener *shortener.Shortener) {
	handlers := newHandlers(shortener)

	group.GET("/links", handlers.paginateHandler)
	group.GET("/links/:id", handlers.retrieveHandler)
	group.POST("/links", handlers.shortenHandler)
	group.PATCH("/links/:id", handlers.updateHandler)
	group.DELETE("/links/:id", handlers.deleteHandler)
}
