package cmd

import (
	"fmt"
	"strconv"

	"github.com/3Danger/currency/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func migrateCmd(ctx context.Context, conf config.Config) *cobra.Command {
	command := &cobra.Command{ //nolint:exhaustruct
		Use:       "migrate",
		Short:     "run db migrations",
		ValidArgs: []string{"postgres"},
		RunE: func(cmd *cobra.Command, args []string) error {
			//nolint:wrapcheck
			return cmd.Usage()
		},
	}

	command.AddCommand(
		postgresCmd(ctx, conf),
	)

	return command
}

type migrationConstructFn func(context.Context, config.Config) (*migrate.Migrate, error)

func up(ctx context.Context, conf config.Config, constructFn migrationConstructFn) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "up",
		Short: "up migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := constructFn(ctx, conf)
			if err != nil {
				return errors.Wrap(err, "construct migration")
			}

			err = m.Up()
			if err != nil {
				if errors.Is(err, migrate.ErrNoChange) || errors.Is(err, migrate.ErrNilVersion) {
					return nil
				}

				return errors.Wrap(err, "up migrations")
			}

			return nil
		},
	}
}

func down(ctx context.Context, conf config.Config, constructFn migrationConstructFn) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "down",
		Short: "rollback all migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := constructFn(ctx, conf)
			if err != nil {
				return errors.Wrap(err, "construct migration")
			}

			err = m.Down()
			if err != nil {
				return errors.Wrap(err, "down migrations")
			}

			return nil
		},
	}
}

func step(ctx context.Context, conf config.Config, constructFn migrationConstructFn) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "step",
		Short: "run N steps migrations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			stepsCount, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.Wrap(err, "steps count must be integer")
			}

			m, err := constructFn(ctx, conf)
			if err != nil {
				return errors.Wrap(err, "construct migration")
			}

			err = m.Steps(stepsCount)
			if err != nil {
				return errors.Wrap(err, "step count migrations")
			}

			return nil
		},
	}
}

func version(ctx context.Context, conf config.Config, constructFn migrationConstructFn) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "version",
		Short: "display current migration version",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := constructFn(ctx, conf)
			if err != nil {
				return errors.Wrap(err, "construct migration")
			}

			ver, dirty, err := m.Version()
			if err != nil {
				return errors.Wrap(err, "display migration version")
			}

			fmt.Printf("current version: %d, dirty: %v", ver, dirty) //nolint:forbidigo

			return nil
		},
	}
}
