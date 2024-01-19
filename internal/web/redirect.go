package web

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/eduardolat/tinylink/internal/database/dbgen"
	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/hashutil"
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/eduardolat/tinylink/internal/web/layouts"
	"github.com/google/uuid"
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

	// Get the values from the form (only in the POST request)
	visitIDStr := c.FormValue("visit_id")
	password := c.FormValue("password")

	// If the visit ID is provided, check if it's valid and use it
	visitID := uuid.NullUUID{
		Valid: false,
	}
	if visitIDStr != "" {
		validUuid, err := uuid.Parse(visitIDStr)
		if err == nil {
			visitID = uuid.NullUUID{
				Valid: true,
				UUID:  validUuid,
			}
		}
	}

	// Retrieve the URL from the shortener service
	link, err := h.shortener.GetByShortCode(shortCode)
	if errors.Is(err, shortener.ErrLinkNotFound) {
		return c.Redirect(http.StatusFound, "/404")
	}
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

	// If we has no visit ID, create a new one
	if !visitID.Valid {
		referer := c.Request().Referer()
		visit, err := h.shortener.CreateVisit(
			link.ID,
			dbgen.Visits_CreateParams{
				LinkID:    link.ID,
				Ip:        c.RealIP(),
				UserAgent: c.Request().UserAgent(),
				Referer: sql.NullString{
					Valid:  referer != "",
					String: referer,
				},
				IsRedirected: false,
			},
		)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		visitID = uuid.NullUUID{
			Valid: true,
			UUID:  visit.ID,
		}
	}

	// If the link is password protected, check if the password
	// is provided and correct. Otherwise, show the password
	// required page.
	if link.Password.Valid && link.Password.String != "" {
		passwordMatch := false
		passwordProvided := password != ""

		if passwordProvided {
			passwordMatch = hashutil.CompareHashAndPassword(link.Password.String, password)
		}

		if !passwordMatch {
			showPasswordError := passwordProvided
			page := redirectPasswordPage(
				shortCode,
				showPasswordError,
				visitID,
			)
			return echoutil.RenderGomponent(c, http.StatusOK, page)
		}
	}

	redirectCode := shortener.HTTPRedirectCodeTemporary
	if link.HttpRedirectCode != 0 {
		redirectCode = int(link.HttpRedirectCode)
	}

	// Mark the visit as redirected
	if visitID.Valid {
		_, err = h.shortener.SetVisitAsRedirected(visitID.UUID)
		if err != nil {
			logger.Error(
				"failed to mark visit as redirected",
				"error",
				err,
			)
		}
	}

	// Redirect the user to the URL
	return c.Redirect(int(redirectCode), link.OriginalUrl)
}

func redirectPasswordPage(
	shortCode string,
	showPasswordError bool,
	visitID uuid.NullUUID,
) gomponents.Node {
	return layouts.Public("Password Required", []gomponents.Node{
		html.H1(gomponents.Text("Password Required")),
		html.FormEl(
			html.Method("POST"),
			html.Action("/"+shortCode),

			gomponents.If(
				visitID.Valid,
				html.Input(
					html.Type("hidden"),
					html.Name("visit_id"),
					html.Value(visitID.UUID.String()),
				),
			),

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
