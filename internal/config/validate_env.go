package config

import "github.com/eduardolat/tinylink/internal/logger"

// validateEnv validates the given environment variables.
// If there is an error, it will log it and exit the program.
func validateEnv(env *Env) {
	if env.DB_TYPE == PostgresDBType {
		validatePostgresEnv(env)
	}
}

// validatePostgresEnv validates the given Postgres environment variables.
// If there is an error, it will log it and exit the program.
func validatePostgresEnv(env *Env) {
	if env.POSTGRES_HOST == "" {
		logger.FatalError("POSTGRES_HOST is required")
	}

	if env.POSTGRES_PORT == 0 {
		logger.FatalError("POSTGRES_PORT is required")
	}

	if env.POSTGRES_USER == "" {
		logger.FatalError("POSTGRES_USER is required")
	}
}
