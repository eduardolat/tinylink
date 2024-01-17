package v1

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func (h *handlers) shorten(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "url query param is required")
	}

	shortCode, err := h.shortener.ShortenURL(shortener.StoreURLParams{
		OriginalURL: url,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, shortCode)
}
