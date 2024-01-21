package admin

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/htmx"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/sqlutil"
	"github.com/eduardolat/tinylink/internal/timeutil"
	"github.com/eduardolat/tinylink/internal/validatorutil"
	"github.com/eduardolat/tinylink/internal/web/components"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents-heroicons/v2/solid"
	"github.com/maragudk/gomponents/html"
)

type paginateRequest struct {
	Page              int      `query:"page" validate:"omitempty,numeric"`
	FilterIsActive    *bool    `query:"filter_is_active" validate:"omitnil,boolean"`
	FilterOriginalUrl *string  `query:"filter_original_url"`
	FilterShortCode   *string  `query:"filter_short_code"`
	FilterDescription *string  `query:"filter_description"`
	FilterTags        []string `query:"filter_tags"`
}

func newPaginateRequest() *paginateRequest {
	return &paginateRequest{
		Page:              1,
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

func (h *handlers) linksPaginateHandler(c echo.Context) error {
	pageSize := 20

	req := newPaginateRequest()
	if err := req.BindAndValidate(c); err != nil {
		return err
	}

	pagination, err := h.shortener.Paginate(shortener.PaginateParams{
		Size:              pageSize,
		Page:              req.Page,
		FilterIsActive:    sqlutil.NullBoolFromPtr(req.FilterIsActive),
		FilterOriginalUrl: sqlutil.NullStringFromPtr(req.FilterOriginalUrl),
		FilterShortCode:   sqlutil.NullStringFromPtr(req.FilterShortCode),
		FilterDescription: sqlutil.NullStringFromPtr(req.FilterDescription),
		FilterTags:        req.FilterTags,
	})
	if err != nil {
		return err
	}

	page := h.linksPaginatePage(pagination)
	return echoutil.RenderGomponent(c, http.StatusOK, page)
}

func (h *handlers) linksPaginatePage(
	pagination shortener.PaginateResponse,
) gomponents.Node {
	nextPageLink := fmt.Sprintf("/admin/links?page=%v", pagination.NextPage)

	nextPageLoader := gomponents.Group([]gomponents.Node{
		htmx.HxGet(nextPageLink),
		htmx.HxTrigger("revealed"),
		htmx.HxSwap("afterend"),
		htmx.HxIndicator("#links-loading"),
	})

	trs := make([]gomponents.Node, len(pagination.Items))
	for i, item := range pagination.Items {
		shortLink := h.shortener.CreateShortLinkFromCode(item.ShortCode)
		isLastTr := i == len(pagination.Items)-1
		hasNextPage := pagination.NextPage != 0

		trs[i] = html.Tr(
			gomponents.If(
				isLastTr && hasNextPage,
				nextPageLoader,
			),

			html.Td(
				components.FlagCircle(components.FlagCircleProps{
					Flag:         item.IsActive,
					TrueTooltip:  "Active",
					FalseTooltip: "Inactive",
				}),
			),
			html.Td(
				html.A(
					html.Href(shortLink),
					html.Target("_blank"),
					html.Class("flex items-center justifys-start"),
					solid.ArrowTopRightOnSquare(html.Class("h-4 w-4 mr-1")),
					html.Span(
						gomponents.Text(
							item.ShortCode,
						),
					),
				),
			),
			html.Td(
				html.Class("truncate max-w-[50px]"),
				html.A(
					html.Href(item.OriginalUrl),
					gomponents.Text(item.OriginalUrl),
				),
			),
			html.Td(
				gomponents.Textf("%v", item.HttpRedirectCode),
			),
			html.Td(
				components.FlagCircle(components.FlagCircleProps{
					Flag:         item.Password.Valid,
					TrueTooltip:  "Password protected",
					FalseTooltip: "Not password protected",
				}),
			),
			html.Td(
				gomponents.Text(timeutil.FormatDateTimeShort(item.CreatedAt)),
			),
		)
	}

	return components.RenderableGroup(trs)
}
