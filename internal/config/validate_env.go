package config

import "github.com/eduardolat/tinylink/internal/logger"

// validateEnv validates the given environment variables.
// If there is an error, it will log it and exit the program.
func validateEnv(env *Env) {
	validatePostgresEnv(env)
	validateBasicAuth(env)
}

// validatePostgresEnv validates the given Postgres environment variables.
// If there is an error, it will log it and exit the program.
func validatePostgresEnv(env *Env) {
	if env.DB_TYPE == nil {
		return
	}

	if *env.DB_TYPE != PostgresDBType {
		return
	}

	if env.POSTGRES_HOST == nil {
		logger.FatalError("POSTGRES_HOST is required")
	}

	if env.POSTGRES_PORT == nil {
		logger.FatalError("POSTGRES_PORT is required")
	}

	if env.POSTGRES_USER == nil {
		logger.FatalError("POSTGRES_USER is required")
	}
}

// validateBasicAuth validates the given basic auth environment variables.
// If there is an error, it will log it and exit the program.
func validateBasicAuth(env *Env) {
	if env.ENABLE_BASIC_AUTH == nil {
		return
	}

	if !*env.ENABLE_BASIC_AUTH {
		return
	}

	if env.BASIC_AUTH_USERNAME == nil {
		logger.FatalError("BASIC_AUTH_USERNAME is required")
	}

	if env.BASIC_AUTH_PASSWORD == nil {
		logger.FatalError("BASIC_AUTH_PASSWORD is required")
	}
}
