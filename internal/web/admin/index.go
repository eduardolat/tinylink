package admin

import "github.com/labstack/echo/v4"

func (h *handlers) indexHandler(c echo.Context) error {
	return c.String(200, "Admin")
}
