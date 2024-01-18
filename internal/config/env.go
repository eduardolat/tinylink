package config

import (
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/joho/godotenv"
)

const (
	InMemoryDBType string = "in-memory"
	PostgresDBType string = "postgres"

	UUIDGeneratorType   string = "uuid"
	NanoIDGeneratorType string = "nanoid"

	DefaultPort          int    = 3000
	DefaultDBType        string = InMemoryDBType
	DefaultGeneratorType string = NanoIDGeneratorType
)

type Env struct {
	// General env variables
	TL_PORT           *int
	TL_DB_TYPE        *string
	TL_GENERATOR_TYPE *string
	TL_URL            *string

	// Basic Auth specific env variables
	TL_ENABLE_BASIC_AUTH   *bool
	TL_BASIC_AUTH_USERNAME *string
	TL_BASIC_AUTH_PASSWORD *string

	// Postgres specific env variables
	TL_POSTGRES_HOST *string
	TL_POSTGRES_PORT *int
	TL_POSTGRES_USER *string
	TL_POSTGRES_PASS *string
	TL_POSTGRES_DB   *string
	TL_POSTGRES_SSL  *bool

	// UUID specific env variables
	TL_UUID_REMOVE_DASHES *bool

	// NanoID specific env variables
	TL_NANOID_SIZE     *int
	TL_NANOID_ALPHABET *string
}

// GetEnv returns the environment variables.
//
// If there is an error, it will log it and exit the program.
func GetEnv() *Env {
	err := godotenv.Load()
	if err == nil {
		logger.Info("👉 using .env file")
	}

	env := &Env{
		// General env variables
		TL_PORT: getEnvAsInt(getEnvAsIntParams{
			name:         "TL_PORT",
			defaultValue: newDefaultValue(DefaultPort),
		}),
		TL_DB_TYPE: getEnvAsString(getEnvAsStringParams{
			name:         "TL_DB_TYPE",
			defaultValue: newDefaultValue(DefaultDBType),
		}),
		TL_GENERATOR_TYPE: getEnvAsString(getEnvAsStringParams{
			name:         "TL_GENERATOR_TYPE",
			defaultValue: newDefaultValue(DefaultGeneratorType),
		}),
		TL_URL: getEnvAsString(getEnvAsStringParams{
			name:       "TL_URL",
			isRequired: true,
		}),

		// Basic Auth specific env variables
		TL_ENABLE_BASIC_AUTH: getEnvAsBool(getEnvAsBoolParams{
			name:       "TL_ENABLE_BASIC_AUTH",
			isRequired: false,
		}),
		TL_BASIC_AUTH_USERNAME: getEnvAsString(getEnvAsStringParams{
			name:       "TL_BASIC_AUTH_USERNAME",
			isRequired: false,
		}),
		TL_BASIC_AUTH_PASSWORD: getEnvAsString(getEnvAsStringParams{
			name:       "TL_BASIC_AUTH_PASSWORD",
			isRequired: false,
		}),

		// Postgres specific env variables
		TL_POSTGRES_HOST: getEnvAsString(getEnvAsStringParams{
			name:       "TL_POSTGRES_HOST",
			isRequired: false,
		}),
		TL_POSTGRES_PORT: getEnvAsInt(getEnvAsIntParams{
			name:       "TL_POSTGRES_PORT",
			isRequired: false,
		}),
		TL_POSTGRES_USER: getEnvAsString(getEnvAsStringParams{
			name:       "TL_POSTGRES_USER",
			isRequired: false,
		}),
		TL_POSTGRES_PASS: getEnvAsString(getEnvAsStringParams{
			name:       "TL_POSTGRES_PASS",
			isRequired: false,
		}),
		TL_POSTGRES_DB: getEnvAsString(getEnvAsStringParams{
			name:       "TL_POSTGRES_DB",
			isRequired: false,
		}),
		TL_POSTGRES_SSL: getEnvAsBool(getEnvAsBoolParams{
			name:       "TL_POSTGRES_ENABLE_SSL",
			isRequired: false,
		}),

		// UUID specific env variables
		TL_UUID_REMOVE_DASHES: getEnvAsBool(getEnvAsBoolParams{
			name:       "TL_UUID_REMOVE_DASHES",
			isRequired: false,
		}),

		// NanoID specific env variables
		TL_NANOID_SIZE: getEnvAsInt(getEnvAsIntParams{
			name:       "TL_NANOID_SIZE",
			isRequired: false,
		}),
		TL_NANOID_ALPHABET: getEnvAsString(getEnvAsStringParams{
			name:       "TL_NANOID_ALPHABET",
			isRequired: false,
		}),
	}

	validateEnv(env)
	return env
}
