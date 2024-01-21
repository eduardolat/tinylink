package admin

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/eduardolat/tinylink/internal/database/dbgen"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/validatorutil"
	"github.com/labstack/echo/v4"
)

type shortenRequest struct {
	ShortCode         string   `form:"short_code"`
	OriginalUrl       string   `form:"original_url" validate:"required,url"`
	HttpRedirectCode  int      `form:"http_redirect_code"`
	IsActive          bool     `form:"is_active"`
	Description       string   `form:"description"`
	Tags              []string `form:"tags"`
	Password          string   `form:"password"`
	ExpiresAt         string   `form:"expires_at" validate:"omitempty,datetime"`
	DuplicateIfExists bool     `form:"duplicate_if_exists"`
}

func newShortenRequest() *shortenRequest {
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
	req := newShortenRequest()
	if err := req.BindAndValidate(c); err != nil {
		return err
	}

	var parsedExpiresAt time.Time
	if req.ExpiresAt != "" {
		var err error
		parsedExpiresAt, err = time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return err
		}
	}

	link, err := h.shortener.Shorten(
		dbgen.Links_CreateParams{
			ShortCode:        req.ShortCode,
			OriginalUrl:      req.OriginalUrl,
			HttpRedirectCode: int16(req.HttpRedirectCode),
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
		return err
	}

	return c.String(http.StatusOK, link.ShortCode)
}
