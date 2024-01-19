package v1

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/eduardolat/tinylink/internal/database/dbgen"
	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/validatorutil"
	"github.com/labstack/echo/v4"
)

type shortenRequest struct {
	ShortCode         string   `json:"short_code"`
	OriginalUrl       string   `json:"original_url" validate:"required,url"`
	HttpRedirectCode  int      `json:"http_redirect_code"`
	Description       string   `json:"description"`
	Tags              []string `json:"tags"`
	Password          string   `json:"password"`
	ExpiresAt         string   `json:"expires_at" validate:"omitempty,datetime"`
	IsActive          bool     `json:"is_active"`
	DuplicateIfExists bool     `json:"duplicate_if_exists"`
}

func NewShortenRequest() *shortenRequest {
	return &shortenRequest{
		ShortCode:         "",
		HttpRedirectCode:  shortener.HTTPRedirectCodeTemporary,
		Tags:              []string{},
		Password:          "",
		ExpiresAt:         "",
		IsActive:          true,
		DuplicateIfExists: false,
	}
}

func (s *shortenRequest) BindAndValidate(c echo.Context) error {
	if err := c.Bind(s); err != nil {
		return err
	}
	return validatorutil.PrettyValidate(s)
}

func (h *handlers) shortenHandler(c echo.Context) error {
	req := NewShortenRequest()
	if err := req.BindAndValidate(c); err != nil {
		return echoutil.JsonError(c, http.StatusBadRequest, err)
	}

	var parsedExpiresAt time.Time
	if req.ExpiresAt != "" {
		var err error
		parsedExpiresAt, err = time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return echoutil.JsonError(c, http.StatusBadRequest, err)
		}
	}

	link, err := h.shortener.ShortenURL(
		dbgen.Links_CreateParams{
			ShortCode:        req.ShortCode,
			OriginalUrl:      req.OriginalUrl,
			HttpRedirectCode: int32(req.HttpRedirectCode),
			Description: sql.NullString{
				Valid:  req.Description != "",
				String: req.Description,
			},
			Tags: req.Tags,
			Password: sql.NullString{
				Valid:  req.Password != "",
				String: req.Password,
			},
			ExpiresAt: sql.NullTime{
				Valid: req.ExpiresAt != "",
				Time:  parsedExpiresAt,
			},
			IsActive: req.IsActive,
			CreatedByIp: sql.NullString{
				Valid:  true,
				String: c.RealIP(),
			},
			CreatedByUserAgent: sql.NullString{
				Valid:  true,
				String: c.Request().UserAgent(),
			},
		},
		req.DuplicateIfExists,
	)
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	shortURL := h.shortener.CreateShortURL(link.ShortCode)

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"short_code": link.ShortCode,
			"short_url":  shortURL,
			"created_at": time.Now().Format(time.RFC3339),
		},
	)
}
