package web

import (
	"net/http"
	"time"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/web/layouts"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) redirectHandler(c echo.Context) error {
	// The short code is the last part of the URL
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		return c.Redirect(http.StatusFound, "/404")
	}

	// Only available in the POST request to check password
	// protected links
	password := c.FormValue("password")

	// Retrieve the URL from the shortener service
	data, err := h.shortener.RetrieveURL(shortCode)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Check if the link is active and not expired
	if !data.IsActive {
		return c.Redirect(http.StatusFound, "/404")
	}
	if data.ExpiresAt.Valid && data.ExpiresAt.Time.Before(time.Now()) {
		return c.Redirect(http.StatusFound, "/404")
	}

	// If the link is password protected, check if the password
	// is provided and correct. Otherwise, show the password
	// required page.
	if data.Password.Valid && data.Password.String != password {
		page := redirectPasswordPage(shortCode)
		return echoutil.RenderGomponent(c, http.StatusOK, page)
	}

	redirectCode := shortener.HTTPRedirectCodeTemporary
	if data.HTTPRedirectCode != 0 {
		redirectCode = data.HTTPRedirectCode
	}

	// Redirect the user to the URL
	return c.Redirect(int(redirectCode), data.OriginalURL)
}

func redirectPasswordPage(shortCode string) gomponents.Node {
	return layouts.Public("Password Required", []gomponents.Node{
		html.H1(gomponents.Text("Password Required")),
		html.FormEl(
			html.Method("POST"),
			html.Action("/"+shortCode),

			html.Label(
				html.For("password"),
				html.Span(
					gomponents.Text("Enter the password to continue"),
				),

				html.Input(
					html.ID("password"),
					html.Type("password"),
					html.Name("password"),
					html.Placeholder("Password"),
				),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text("Submit"),
			),
		),
	})
}
