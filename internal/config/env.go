package config

import (
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/joho/godotenv"
)

const (
	GeneratorTypeUUID   string = "uuid"
	GeneratorTypeNanoID string = "nanoid"

	DefaultPort          int    = 3000
	DefaultGeneratorType string = GeneratorTypeNanoID
)

type Env struct {
	// General env variables
	TL_PORT           *int
	TL_GENERATOR_TYPE *string
	TL_BASE_URL       *string

	// Postgres specific env variables
	TL_POSTGRES_CONNECTION_STRING *string

	// Basic Auth specific env variables
	TL_ENABLE_BASIC_AUTH   *bool
	TL_BASIC_AUTH_USERNAME *string
	TL_BASIC_AUTH_PASSWORD *string

	// NanoID specific env variables
	TL_NANOID_SIZE     *int
	TL_NANOID_ALPHABET *string

	// UUID specific env variables
	TL_UUID_REMOVE_DASHES *bool
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
		TL_GENERATOR_TYPE: getEnvAsString(getEnvAsStringParams{
			name:         "TL_GENERATOR_TYPE",
			defaultValue: newDefaultValue(DefaultGeneratorType),
		}),
		TL_BASE_URL: getEnvAsString(getEnvAsStringParams{
			name:       "TL_BASE_URL",
			isRequired: true,
		}),

		// Postgres specific env variables
		TL_POSTGRES_CONNECTION_STRING: getEnvAsString(getEnvAsStringParams{
			name:       "TL_POSTGRES_CONNECTION_STRING",
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

		// NanoID specific env variables
		TL_NANOID_SIZE: getEnvAsInt(getEnvAsIntParams{
			name:       "TL_NANOID_SIZE",
			isRequired: false,
		}),
		TL_NANOID_ALPHABET: getEnvAsString(getEnvAsStringParams{
			name:       "TL_NANOID_ALPHABET",
			isRequired: false,
		}),

		// UUID specific env variables
		TL_UUID_REMOVE_DASHES: getEnvAsBool(getEnvAsBoolParams{
			name:       "TL_UUID_REMOVE_DASHES",
			isRequired: false,
		}),
	}

	validateEnv(env)
	return env
}
