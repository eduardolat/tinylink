package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/eduardolat/tinylink/internal/config"
	"github.com/jackc/pgx/v5"
)

// DataStore is a postgres implementation of the DataStore interface
type DataStore struct {
	conn *pgx.Conn
}

// NewDataStore returns a new postgres DataStore and ensures the connection is
// valid and working
func NewDataStore(env *config.Env) (*DataStore, error) {
	if env.TL_POSTGRES_CONNECTION_STRING == nil {
		return nil, errors.New("missing postgres connection string")
	}

	conn, err := pgx.Connect(
		context.Background(),
		*env.TL_POSTGRES_CONNECTION_STRING,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres data store: %w", err)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres data store: %w", err)
	}

	return &DataStore{
		conn: conn,
	}, nil
}
