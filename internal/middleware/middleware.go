package middleware

import "github.com/eduardolat/tinylink/internal/config"

// Middleware is the struct that holds all the middlewares for the
// application.
type Middleware struct {
	env *config.Env
}

// NewMiddleware returns a new instance of the Middleware struct.
func NewMiddleware(env *config.Env) *Middleware {
	return &Middleware{
		env: env,
	}
}
