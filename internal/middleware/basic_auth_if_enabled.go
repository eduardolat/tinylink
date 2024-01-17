package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const basicAuthRealm string = "Please Login"

// BasicAuthIfEnabled is a middleware that checks if the basic auth is enabled
// in the application configuration and if it is, it will check the credentials
// sent in the request.
func (mid *Middleware) BasicAuthIfEnabled(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if mid.env.ENABLE_BASIC_AUTH == nil ||
			mid.env.BASIC_AUTH_USERNAME == nil ||
			mid.env.BASIC_AUTH_PASSWORD == nil {
			return next(c)
		}

		if !*mid.env.ENABLE_BASIC_AUTH {
			return next(c)
		}

		username, password, ok := c.Request().BasicAuth()
		if ok {
			if username == *mid.env.BASIC_AUTH_USERNAME &&
				password == *mid.env.BASIC_AUTH_PASSWORD {
				return next(c)
			}
		}

		c.Response().Header().Set(
			"WWW-Authenticate",
			`Basic realm="`+basicAuthRealm+`", charset="UTF-8"`,
		)
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
}
