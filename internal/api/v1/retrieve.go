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

	link, err := h.shortener.GetByShortCode(shortCode)
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	shortUrl := h.shortener.CreateShortLinkFromCode(link.ShortCode)

	response := retrieveResponse{
		ShortURL:         shortUrl,
		IsActive:         link.IsActive,
		Tags:             link.Tags,
		HTTPRedirectCode: int(link.HttpRedirectCode),
		OriginalURL:      link.OriginalUrl,
		ShortCode:        link.ShortCode,
		CreatedAt:        link.CreatedAt,
		UpdatedAt:        link.UpdatedAt,
	}

	if link.Description.Valid {
		response.Description = &link.Description.String
	}

	if link.ExpiresAt.Valid {
		response.ExpiresAt = &link.ExpiresAt.Time
	}

	if link.CreatedByIp.Valid {
		response.CreatedByIP = &link.CreatedByIp.String
	}

	if link.CreatedByUserAgent.Valid {
		response.CreatedByUserAgent = &link.CreatedByUserAgent.String
	}

	return c.JSON(http.StatusOK, response)
}
