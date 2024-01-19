package config

import "github.com/eduardolat/tinylink/internal/logger"

// validateEnv validates the given environment variables.
// If there is an error, it will log it and exit the program.
func validateEnv(env *Env) {
	validateBasicAuth(env)
}

// validateBasicAuth validates the given basic auth environment variables.
// If there is an error, it will log it and exit the program.
func validateBasicAuth(env *Env) {
	if env.TL_ENABLE_BASIC_AUTH == nil {
		return
	}

	if !*env.TL_ENABLE_BASIC_AUTH {
		return
	}

	if env.TL_BASIC_AUTH_USERNAME == nil {
		logger.FatalError("TL_BASIC_AUTH_USERNAME is required")
	}

	if env.TL_BASIC_AUTH_PASSWORD == nil {
		logger.FatalError("TL_BASIC_AUTH_PASSWORD is required")
	}
}
