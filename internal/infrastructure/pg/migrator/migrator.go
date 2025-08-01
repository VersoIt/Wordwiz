package migrator

import "github.com/jmoiron/sqlx"

const (
	pg            = "postgres"
	migrationsDir = "./migrations"
)

type Migrator struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}
