package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/labstack/echo/v4"
)

type retrieveResponse struct {
	ShortURL           string     `json:"short_url"`
	IsActive           bool       `json:"is_active"`
	Description        *string    `json:"description"`
	Tags               []string   `json:"tags"`
	HTTPRedirectCode   int        `json:"http_redirect_code"`
	OriginalURL        string     `json:"original_url"`
	ShortCode          string     `json:"short_code"`
	Clicks             *int64     `json:"clicks"`
	FirstClickAt       *time.Time `json:"first_click_at"`
	LastClickAt        *time.Time `json:"last_click_at"`
	Redirects          *int64     `json:"redirects"`
	FirstRedirectAt    *time.Time `json:"first_redirect_at"`
	LastRedirectAt     *time.Time `json:"last_redirect_at"`
	ExpiresAt          *time.Time `json:"expires_at"`
	CreatedByIP        *string    `json:"created_by_ip"`
	CreatedByUserAgent *string    `json:"created_by_user_agent"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

func (h *handlers) retrieveHandler(c echo.Context) error {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		return echoutil.JsonError(c, http.StatusBadRequest, errors.New("short code is required"))
	}

	url, err := h.shortener.RetrieveURL(shortCode)
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	shortUrl := h.shortener.CreateShortURL(url.ShortCode)

	response := retrieveResponse{
		ShortURL:         shortUrl,
		IsActive:         url.IsActive,
		Tags:             url.Tags,
		HTTPRedirectCode: url.HTTPRedirectCode,
		OriginalURL:      url.OriginalURL,
		ShortCode:        url.ShortCode,
		CreatedAt:        url.CreatedAt,
		UpdatedAt:        url.UpdatedAt,
	}

	if url.Description.Valid {
		response.Description = &url.Description.String
	}

	if url.Clicks.Valid {
		response.Clicks = &url.Clicks.Int64
	}

	if url.FirstClickAt.Valid {
		response.FirstClickAt = &url.FirstClickAt.Time
	}

	if url.LastClickAt.Valid {
		response.LastClickAt = &url.LastClickAt.Time
	}

	if url.Redirects.Valid {
		response.Redirects = &url.Redirects.Int64
	}

	if url.FirstRedirectAt.Valid {
		response.FirstRedirectAt = &url.FirstRedirectAt.Time
	}

	if url.LastRedirectAt.Valid {
		response.LastRedirectAt = &url.LastRedirectAt.Time
	}

	if url.ExpiresAt.Valid {
		response.ExpiresAt = &url.ExpiresAt.Time
	}

	if url.CreatedByIP.Valid {
		response.CreatedByIP = &url.CreatedByIP.String
	}

	if url.CreatedByUserAgent.Valid {
		response.CreatedByUserAgent = &url.CreatedByUserAgent.String
	}

	return c.JSON(http.StatusOK, response)
}
