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

// TODO: Insert the visits and mark the link as visited
// TODO: Hash the password before storing it in the database and before comparing

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
	link, err := h.shortener.GetByShortCode(shortCode)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Check if the link is active and not expired
	if !link.IsActive {
		return c.Redirect(http.StatusFound, "/404")
	}
	if link.ExpiresAt.Valid && link.ExpiresAt.Time.Before(time.Now()) {
		return c.Redirect(http.StatusFound, "/404")
	}

	// If the link is password protected, check if the password
	// is provided and correct. Otherwise, show the password
	// required page.
	if link.Password.Valid && link.Password.String != password {
		showPasswordError := password != ""
		page := redirectPasswordPage(shortCode, showPasswordError)
		return echoutil.RenderGomponent(c, http.StatusOK, page)
	}

	redirectCode := shortener.HTTPRedirectCodeTemporary
	if link.HttpRedirectCode != 0 {
		redirectCode = int(link.HttpRedirectCode)
	}

	// Redirect the user to the URL
	return c.Redirect(int(redirectCode), link.OriginalUrl)
}

func redirectPasswordPage(shortCode string, showPasswordError bool) gomponents.Node {
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
					html.Required(),
					html.AutoFocus(),
					gomponents.If(
						showPasswordError,
						html.Aria("invalid", "true"),
					),
				),
				gomponents.If(
					showPasswordError,
					html.Small(
						gomponents.Text("Incorrect password"),
					),
				),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text("Submit"),
			),
		),
	})
}
