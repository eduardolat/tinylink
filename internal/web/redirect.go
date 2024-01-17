package web

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/labstack/echo/v4"
)

func (h *handlers) redirectHandler(c echo.Context) error {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		return c.String(http.StatusBadRequest, "shortCode is required")
	}

	// Retrieve the URL from the shortener service
	data, err := h.shortener.RetrieveURL(shortCode)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	redirectCode := shortener.HTTPRedirectCodeTemporary
	if data.HTTPRedirectCode != 0 {
		redirectCode = data.HTTPRedirectCode
	}

	// Redirect the user to the URL
	return c.Redirect(int(redirectCode), data.OriginalURL)
}
