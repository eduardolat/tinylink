package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) notFoundHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "Not Found")
}
