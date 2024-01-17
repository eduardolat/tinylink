package admin

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/web/layouts"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexHandler(c echo.Context) error {
	page := indexPage()
	return echoutil.RenderGomponent(c, http.StatusOK, page)
}

func indexPage() gomponents.Node {
	return layouts.Admin("Admin", []gomponents.Node{
		html.H1(gomponents.Text("Admin")),
	})
}
