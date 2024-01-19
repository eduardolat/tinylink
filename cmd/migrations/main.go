package main

/*
	This command uses the goose library under the hood to create new
	migration files and down the latest migration.

	1. Create a new migration file:

	This will create a new migration file in the migrations folder
	Use one of the following commands:
	- go run ./cmd/migrations/. create <migration_name>
	- task migrations:create -- <migration_name>


	2. Up the migrations:

	This will up ALL the migrations in the migrations folder.

	Use one of the following commands:
	- go run ./cmd/migrations/. up
	- task migrations:up

	3. Down the latest migration:

	This roll back the version by 1

	Use one of the following commands:
	- go run ./cmd/migrations/. down
	- task migrations:down

	4. Reset the migrations:

	This will down all the migrations.

	Use one of the following commands:
	- go run ./cmd/migrations/. reset
	- task migrations:reset

	++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	⚠️ TinyLink startup migrations ⚠️

	The migrations up's automatically when the main server starts, so we don't
	need to run it manually with this command.

	The reason to run the up's automatically is to avoid the need to
	manually run the migrations when deploying to production so we can
	provide new versions of the app without the user having to worry
	about running the migrations.

	All the migrations should not be destructive, so we can run them
	without worrying about losing data.
*/

import (
	"context"
	"database/sql"
	"os"

	"github.com/eduardolat/tinylink/internal/config"
	"github.com/eduardolat/tinylink/internal/database"
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/pressly/goose/v3"
)

func main() {
	if len(os.Args) < 2 {
		logger.FatalError(
			"you must provide one of the following commands: create, up, down, reset",
		)
	}
	command := os.Args[1]

	if command != "create" && command != "up" && command != "down" && command != "reset" {
		logger.FatalError(
			"you must provide one of the following commands: create, up, down, reset",
		)
	}

	if command == "create" && len(os.Args) < 3 {
		logger.FatalError(
			"you must provide the name of the migration",
		)
	}

	env := config.GetEnv()
	db, err := database.Connect(env)
	if err != nil {
		logger.FatalError(
			"could not connect to DB",
			"error",
			err,
		)
	}

	if command == "create" {
		migrationName := os.Args[2]
		err := createMigration(migrationName)
		if err != nil {
			logger.FatalError(
				"could not create migration",
				"error",
				err,
			)
		}
		return
	}

	if command == "up" {
		err := upMigration(db)
		if err != nil {
			logger.FatalError(
				"could not up migration",
				"error",
				err,
			)
		}
	}

	if command == "down" {
		err := downMigration(db)
		if err != nil {
			logger.FatalError(
				"could not down migration",
				"error",
				err,
			)
		}
	}

	if command == "reset" {
		err := resetMigration(db)
		if err != nil {
			logger.FatalError(
				"could not reset migration",
				"error",
				err,
			)
		}
	}
}

func createMigration(name string) error {
	return goose.RunContext(
		context.Background(),
		"create",
		nil,
		"./internal/database/migrations",
		name,
		"sql",
	)
}

func upMigration(db *sql.DB) error {
	return goose.RunContext(
		context.Background(),
		"up",
		db,
		"./internal/database/migrations",
	)
}

func downMigration(db *sql.DB) error {
	return goose.RunContext(
		context.Background(),
		"down",
		db,
		"./internal/database/migrations",
	)
}

func resetMigration(db *sql.DB) error {
	return goose.RunContext(
		context.Background(),
		"reset",
		db,
		"./internal/database/migrations",
	)
}
