package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/eduardolat/tinylink/internal/database/dbgen"
	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/sqlutil"
	"github.com/eduardolat/tinylink/internal/validatorutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type updateRequest struct {
	ShortCode        *string   `json:"short_code"`
	OriginalUrl      *string   `json:"original_url" validate:"omitnil,url"`
	HttpRedirectCode *int      `json:"http_redirect_code"`
	IsActive         *bool     `json:"is_active"`
	Description      *string   `json:"description"`
	Tags             *[]string `json:"tags"`
	Password         *string   `json:"password"`
	ExpiresAt        *string   `json:"expires_at" validate:"omitnil,datetime"`
}

func newUpdateRequest() *updateRequest {
	return &updateRequest{}
}

func (s *updateRequest) BindAndValidate(c echo.Context) error {
	if err := c.Bind(s); err != nil {
		return err
	}
	return validatorutil.PrettyValidate(s)
}

func (h *handlers) updateHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return echoutil.JsonError(c, http.StatusBadRequest, err)
	}

	req := newUpdateRequest()
	if err := req.BindAndValidate(c); err != nil {
		return echoutil.JsonError(c, http.StatusBadRequest, err)
	}

	var parsedExpiresAt sql.NullTime
	if req.ExpiresAt != nil {
		var err error
		parsed, err := time.Parse(time.RFC3339, *req.ExpiresAt)
		if err != nil {
			return echoutil.JsonError(c, http.StatusBadRequest, err)
		}
		parsedExpiresAt = sql.NullTime{
			Valid: true,
			Time:  parsed,
		}
	}

	httpRedirectCode := sql.NullInt16{
		Valid: false,
	}
	if req.HttpRedirectCode != nil {
		httpRedirectCode = sql.NullInt16{
			Valid: true,
			Int16: int16(*req.HttpRedirectCode),
		}
	}

	currentLink, err := h.shortener.Get(id)
	if errors.Is(err, shortener.ErrLinkNotFound) {
		return echoutil.JsonError(c, http.StatusNotFound, err)
	}
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	updatedTags := currentLink.Tags
	if req.Tags != nil {
		updatedTags = *req.Tags
	}

	link, err := h.shortener.Update(
		id,
		dbgen.Links_UpdateParams{
			ShortCode:        sqlutil.NullStringFromPtr(req.ShortCode),
			OriginalUrl:      sqlutil.NullStringFromPtr(req.OriginalUrl),
			HttpRedirectCode: httpRedirectCode,
			Description:      sqlutil.NullStringFromPtr(req.Description),
			Tags:             updatedTags,
			Password:         sqlutil.NullStringFromPtr(req.Password),
			ExpiresAt:        parsedExpiresAt,
			IsActive:         sqlutil.NullBoolFromPtr(req.IsActive),
		},
	)
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(
		http.StatusOK,
		h.linkToJSON(link),
	)
}
