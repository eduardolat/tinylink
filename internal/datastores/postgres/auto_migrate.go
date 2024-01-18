package postgres

import (
	"context"
	"fmt"

	"github.com/eduardolat/tinylink/internal/config"
	"github.com/eduardolat/tinylink/internal/logger"
)

// AutoMigrate runs all migrations that haven't been run yet
// and returns an error if any of them fail
func (ds *DataStore) AutoMigrate() error {
	if err := ds.checkAndInitDB(); err != nil {
		return fmt.Errorf("automigrate: %w", err)
	}

	if err := ds.runMigration01(); err != nil {
		return fmt.Errorf("automigrate: %w", err)
	}

	return nil
}

// getLastMigrationNumber returns the last migration number
// from the tl_migrations table
func (ds *DataStore) getLastMigrationNumber() (int, error) {
	sqlQuery := `
		SELECT migration_number
		FROM tl_migrations
		ORDER BY id DESC
		LIMIT 1;
	`
	var migrationNumber int
	err := ds.conn.QueryRow(
		context.Background(),
		sqlQuery,
	).Scan(&migrationNumber)
	if err != nil {
		return 0, fmt.Errorf("failed to get last migration number: %w", err)
	}

	return migrationNumber, nil
}

// checkAndInitDB checks if the database exists, then creates
// the migrations table if it doesn't exist, then returns the
// last migration number
func (ds *DataStore) checkAndInitDB() error {
	searchPath := ds.conn.Config().RuntimeParams["search_path"]
	if searchPath == "" {
		searchPath = "public"
	}

	sqlQuery := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = $1
			AND table_name = 'tl_migrations'
		);
	`
	var exists bool
	err := ds.conn.QueryRow(
		context.Background(),
		sqlQuery,
		searchPath,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if migrations table exists: %w", err)
	}

	if !exists {
		sqlQuery = `
			CREATE TABLE tl_migrations (
				migration_number INTEGER NOT NULL PRIMARY KEY,
				created_at TIMESTAMP NOT NULL DEFAULT NOW()
			);
		`
		_, err = ds.conn.Exec(
			context.Background(),
			sqlQuery,
		)
		if err != nil {
			return fmt.Errorf("failed to create migrations table: %w", err)
		}

		sqlQuery = `
			INSERT INTO tl_migrations (migration_number)
			VALUES (0);
		`
		_, err = ds.conn.Exec(
			context.Background(),
			sqlQuery,
		)
		if err != nil {
			return fmt.Errorf("failed to insert first migration number: %w", err)
		}

		logger.Info(
			"migrations table created",
			"db_type", config.PostgresDBType,
		)
	}

	return nil
}

// runMigration01 creates the tl_links table
func (ds *DataStore) runMigration01() error {
	lastMigrationNumber, err := ds.getLastMigrationNumber()
	if err != nil {
		return fmt.Errorf("migration01: %w", err)
	}

	if lastMigrationNumber >= 1 {
		return nil
	}

	sqlQuery := `
		CREATE TABLE IF NOT EXISTS links (
			short_code TEXT NOT NULL PRIMARY KEY,
			original_url TEXT NOT NULL,
			http_redirect_code INTEGER NOT NULL,
			is_active BOOLEAN NOT NULL,
			description TEXT,
			tags TEXT[],
			password TEXT,
			clicks INTEGER,
			first_click_at TIMESTAMP,
			last_click_at TIMESTAMP,
			redirects INTEGER,
			first_redirect_at TIMESTAMP,
			last_redirect_at TIMESTAMP,
			expires_at TIMESTAMP,
			created_by_ip TEXT,
			created_by_user_agent TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);

		CREATE INDEX IF NOT EXISTS links_short_code_idx ON links (short_code);
		CREATE INDEX IF NOT EXISTS links_original_url_idx ON links (original_url);
		CREATE INDEX IF NOT EXISTS links_is_active_idx ON links (is_active);
		CREATE INDEX IF NOT EXISTS links_tags_idx ON links USING GIN (tags);
	`
	_, err = ds.conn.Exec(
		context.Background(),
		sqlQuery,
	)
	if err != nil {
		return fmt.Errorf("migration01: %w", err)
	}

	logger.Info(
		"run migration",
		"db_type", config.PostgresDBType,
		"number", 1,
	)

	return nil
}
