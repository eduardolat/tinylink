package config

import (
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/joho/godotenv"
)

type DBType string
type GeneratorType string

const (
	InMemoryDBType DBType = "in-memory"
	PostgresDBType DBType = "postgres"

	UUIDGeneratorType   GeneratorType = "uuid"
	NanoIDGeneratorType GeneratorType = "nanoid"

	DefaultPort          int           = 3000
	DefaultDBType        DBType        = InMemoryDBType
	DefaultGeneratorType GeneratorType = NanoIDGeneratorType
)

type Env struct {
	// General env variables
	PORT           int
	DB_TYPE        DBType
	GENERATOR_TYPE GeneratorType

	// Basic Auth specific env variables
	ENABLE_BASIC_AUTH   bool
	BASIC_AUTH_USERNAME string
	BASIC_AUTH_PASSWORD string

	// Postgres specific env variables
	POSTGRES_HOST string
	POSTGRES_PORT int
	POSTGRES_USER string
	POSTGRES_PASS string
	POSTGRES_DB   string
	POSTGRES_SSL  bool

	// UUID specific env variables
	UUID_REMOVE_DASHES bool

	// NanoID specific env variables
	NANOID_SIZE     int
	NANOID_ALPHABET string
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
		PORT: getEnvAsInt(getEnvAsIntParams{
			name:         "PORT",
			defaultValue: newDefaultValue(DefaultPort),
		}),
		DB_TYPE: DBType(getEnvAsString(getEnvAsStringParams{
			name:         "DB_TYPE",
			defaultValue: newDefaultValue(string(DefaultDBType)),
		})),
		GENERATOR_TYPE: GeneratorType(getEnvAsString(getEnvAsStringParams{
			name:         "GENERATOR_TYPE",
			defaultValue: newDefaultValue(string(DefaultGeneratorType)),
		})),

		// Basic Auth specific env variables
		ENABLE_BASIC_AUTH: getEnvAsBool(getEnvAsBoolParams{
			name:       "ENABLE_BASIC_AUTH",
			isRequired: false,
		}),
		BASIC_AUTH_USERNAME: getEnvAsString(getEnvAsStringParams{
			name:       "BASIC_AUTH_USERNAME",
			isRequired: false,
		}),
		BASIC_AUTH_PASSWORD: getEnvAsString(getEnvAsStringParams{
			name:       "BASIC_AUTH_PASSWORD",
			isRequired: false,
		}),

		// Postgres specific env variables
		POSTGRES_HOST: getEnvAsString(getEnvAsStringParams{
			name:       "POSTGRES_HOST",
			isRequired: false,
		}),
		POSTGRES_PORT: getEnvAsInt(getEnvAsIntParams{
			name:       "POSTGRES_PORT",
			isRequired: false,
		}),
		POSTGRES_USER: getEnvAsString(getEnvAsStringParams{
			name:       "POSTGRES_USER",
			isRequired: false,
		}),
		POSTGRES_PASS: getEnvAsString(getEnvAsStringParams{
			name:       "POSTGRES_PASS",
			isRequired: false,
		}),
		POSTGRES_DB: getEnvAsString(getEnvAsStringParams{
			name:       "POSTGRES_DB",
			isRequired: false,
		}),
		POSTGRES_SSL: getEnvAsBool(getEnvAsBoolParams{
			name:       "POSTGRES_SSL",
			isRequired: false,
		}),

		// UUID specific env variables
		UUID_REMOVE_DASHES: getEnvAsBool(getEnvAsBoolParams{
			name:       "UUID_REMOVE_DASHES",
			isRequired: false,
		}),

		// NanoID specific env variables
		NANOID_SIZE: getEnvAsInt(getEnvAsIntParams{
			name:       "NANOID_SIZE",
			isRequired: false,
		}),
		NANOID_ALPHABET: getEnvAsString(getEnvAsStringParams{
			name:       "NANOID_ALPHABET",
			isRequired: false,
		}),
	}

	validateEnv(env)
	return env
}
