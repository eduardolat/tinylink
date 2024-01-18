package web

import (
	"github.com/eduardolat/tinylink/internal/middleware"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/web/admin"
	"github.com/labstack/echo/v4"
)

func MountRouter(
	group *echo.Group,
	mid *middleware.Middleware,
	shortener *shortener.Shortener,
) {
	handlers := newHandlers(shortener)

	group.Any("/", handlers.notFoundHandler)
	group.GET("/404", handlers.notFoundHandler)

	group.GET("/:shortCode", handlers.redirectHandler)
	group.POST("/:shortCode", handlers.redirectHandler)

	adminGroup := group.Group("/admin")
	admin.MountRouter(adminGroup, mid, shortener)
}
