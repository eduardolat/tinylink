package v1

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/sqlutil"
	"github.com/eduardolat/tinylink/internal/validatorutil"
	"github.com/labstack/echo/v4"
)

type paginateRequest struct {
	Page              int      `query:"page" validate:"omitempty,numeric"`
	Size              int      `query:"size" validate:"omitempty,numeric"`
	FilterIsActive    *bool    `query:"filter_is_active" validate:"omitnil,boolean"`
	FilterOriginalUrl *string  `query:"filter_original_url"`
	FilterShortCode   *string  `query:"filter_short_code"`
	FilterDescription *string  `query:"filter_description"`
	FilterTags        []string `query:"filter_tags"`
}

func newPaginateRequest() *paginateRequest {
	return &paginateRequest{
		Page:              1,
		Size:              10,
		FilterIsActive:    nil,
		FilterOriginalUrl: nil,
		FilterShortCode:   nil,
		FilterDescription: nil,
		FilterTags:        []string{},
	}
}

func (s *paginateRequest) BindAndValidate(c echo.Context) error {
	if err := c.Bind(s); err != nil {
		return err
	}
	return validatorutil.PrettyValidate(s)
}

type paginateResponse struct {
	Page       int        `json:"page"`
	PrevPage   int        `json:"prev_page"`
	NextPage   int        `json:"next_page"`
	TotalPages int        `json:"total_pages"`
	TotalItems int        `json:"total_items"`
	Items      []jsonLink `json:"items"`
}

func (h *handlers) paginateHandler(c echo.Context) error {
	req := newPaginateRequest()
	if err := req.BindAndValidate(c); err != nil {
		return echoutil.JsonError(c, http.StatusBadRequest, err)
	}

	paginateResp, err := h.shortener.Paginate(shortener.PaginateParams{
		Page:              req.Page,
		Size:              req.Size,
		FilterIsActive:    sqlutil.NullBoolFromPtr(req.FilterIsActive),
		FilterOriginalUrl: sqlutil.NullStringFromPtr(req.FilterOriginalUrl),
		FilterShortCode:   sqlutil.NullStringFromPtr(req.FilterShortCode),
		FilterDescription: sqlutil.NullStringFromPtr(req.FilterDescription),
		FilterTags:        req.FilterTags,
	})
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	items := make([]jsonLink, len(paginateResp.Items))
	for i, item := range paginateResp.Items {
		items[i] = h.linkToJSON(item)
	}

	endpointResp := paginateResponse{
		Page:       paginateResp.Page,
		PrevPage:   paginateResp.PrevPage,
		NextPage:   paginateResp.NextPage,
		TotalPages: paginateResp.TotalPages,
		TotalItems: paginateResp.TotalItems,
		Items:      items,
	}

	return c.JSON(http.StatusOK, endpointResp)
}
