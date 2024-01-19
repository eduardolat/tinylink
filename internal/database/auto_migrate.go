package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// AutoMigrate runs goose migrations on the database.
//
// ⚠️ TinyLink startup migrations ⚠️
//
// The migrations up's automatically when the main server starts, so we don't
// need to run it manually with this command.
//
// The reason to run the up's automatically is to avoid the need to
// manually run the migrations when deploying to production so we can
// provide new versions of the app without the user having to worry
// about running the migrations.
//
// All the migrations should not be destructive, so we can run them
// without worrying about losing data.
func AutoMigrate(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed to run goose migrations: %w", err)
	}

	return nil
}
