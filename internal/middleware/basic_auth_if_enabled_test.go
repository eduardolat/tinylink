package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardolat/tinylink/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuthIfEnabled_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up middleware with basic auth enabled and known credentials
	mid := &Middleware{
		env: &config.Env{
			ENABLE_BASIC_AUTH:   pointerToBool(true),
			BASIC_AUTH_USERNAME: pointerToString("username"),
			BASIC_AUTH_PASSWORD: pointerToString("password"),
		},
	}

	// Set up request with correct basic auth credentials
	req.SetBasicAuth("username", "password")

	// Call middleware
	mid.BasicAuthIfEnabled(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)

	// Assert that the middleware allowed the request to proceed
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}

func TestBasicAuthIfEnabled_IncorrectCredentials(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up middleware with basic auth enabled and known credentials
	mid := &Middleware{
		env: &config.Env{
			ENABLE_BASIC_AUTH:   pointerToBool(true),
			BASIC_AUTH_USERNAME: pointerToString("username"),
			BASIC_AUTH_PASSWORD: pointerToString("password"),
		},
	}

	// Set up request with incorrect basic auth credentials
	req.SetBasicAuth("wrongusername", "wrongpassword")

	// Call middleware
	mid.BasicAuthIfEnabled(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)

	// Assert that the middleware rejected the request
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	// Check the www-authenticate header
	assert.NotEmpty(t, rec.Header().Get("WWW-Authenticate"))
	assert.Contains(t, rec.Header().Get("WWW-Authenticate"), basicAuthRealm)
}

func TestBasicAuthIfEnabled_Disabled(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up middleware with basic auth disabled
	mid := &Middleware{
		env: &config.Env{
			ENABLE_BASIC_AUTH:   pointerToBool(false),
			BASIC_AUTH_USERNAME: pointerToString("username"),
			BASIC_AUTH_PASSWORD: pointerToString("password"),
		},
	}

	// Call middleware
	mid.BasicAuthIfEnabled(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)

	// Assert that the middleware allowed the request to proceed
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}

// Helper functions to create pointers to bool and string
func pointerToBool(b bool) *bool {
	return &b
}

func pointerToString(s string) *string {
	return &s
}
