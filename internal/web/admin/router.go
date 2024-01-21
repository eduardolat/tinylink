package admin

import (
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
	handlers := newHandlers(shortener)

	group.GET("", handlers.indexHandler)
	group.GET("/links/paginate", handlers.linksPaginateHandler)
}
