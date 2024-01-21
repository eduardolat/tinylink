package admin

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/htmx"
	"github.com/eduardolat/tinylink/internal/web/components"
	"github.com/eduardolat/tinylink/internal/web/layouts"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexHandler(c echo.Context) error {
	linksCount, err := h.shortener.CountAll()
	if err != nil {
		return err
	}

	page := h.indexPage(linksCount)
	return echoutil.RenderGomponent(c, http.StatusOK, page)
}

func (h *handlers) indexPage(linksCount int64) gomponents.Node {
	noLinksIndicator := html.Div(
		components.Classes{
			"w-full flex justify-center items-center h-30": true,
		},
		html.Strong(
			gomponents.Text("No links yet ✂️"),
		),
	)

	linksTable := html.Div(
		html.Class("overflow-x-auto overflow-y-hidden w-full"),

		html.Table(
			html.THead(
				html.Tr(
					html.Th(
						gomponents.Text("Active?"),
					),
					html.Th(
						gomponents.Text("Short link"),
					),
					html.Th(
						gomponents.Text("Original link"),
					),
					html.Th(
						gomponents.Text("Redirect code"),
					),
					html.Th(
						gomponents.Text("Password?"),
					),
					html.Th(
						gomponents.Text("Created at"),
					),
				),
			),
			html.TBody(
				htmx.HxGet("/admin/links/paginate?page=1"),
				htmx.HxTrigger("load"),
				htmx.HxIndicator("#links-loading"),
			),
		),

		components.HxLoading(components.HxLoadingProps{
			ID:     "links-loading",
			Center: true,
			Size:   "lg",
		}),
	)

	return layouts.Admin("Admin", []gomponents.Node{
		gomponents.If(
			linksCount == 0,
			noLinksIndicator,
		),

		gomponents.If(
			linksCount > 0,
			linksTable,
		),
	})
}
