package migrator

import "github.com/pressly/goose/v3"

func (m *Migrator) Migrate() error {
	err := goose.SetDialect(pg)
	if err != nil {
		return err
	}

	err = goose.Up(m.db.DB, migrationsDir)
	if err != nil {
		return err
	}

	return nil
}
