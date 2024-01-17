package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) retrieveHandler(c echo.Context) error {
	shortCode := c.QueryParam("shortCode")
	if shortCode == "" {
		return c.String(http.StatusBadRequest, "shortCode query param is required")
	}

	url, err := h.shortener.RetrieveURL(shortCode)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, url.OriginalURL)
}
