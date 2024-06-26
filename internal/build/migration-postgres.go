package build

import (
	migration "github.com/3Danger/currency/internal/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // driver
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
)

func (b *Builder) PostgresMigration() (*migrate.Migrate, error) {
	d, err := iofs.New(migration.FS, migration.PostgresPath)
	if err != nil {
		return nil, errors.Wrap(err, "embed postgres migrations")
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, b.cnf.Postgres.DSN())
	if err != nil {
		return nil, errors.Wrap(err, "apply postgres migrations")
	}

	return m, nil
}
