package database

import (
	"database/sql"
	"fmt"

	"github.com/eduardolat/tinylink/internal/config"
	"github.com/eduardolat/tinylink/internal/logger"
	_ "github.com/lib/pq"
)

func Connect(env *config.Env) (*sql.DB, error) {
	if env.TL_POSTGRES_CONNECTION_STRING == nil {
		return nil, fmt.Errorf("the postgres connection string is required")
	}

	db, err := sql.Open("postgres", *env.TL_POSTGRES_CONNECTION_STRING)
	if err != nil {
		return nil, fmt.Errorf("could not connect to DB: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping DB: %w", err)
	}

	logger.Info("✅ connected to DB")

	return db, nil
}
