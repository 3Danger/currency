package cmd

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/build"
	"github.com/3Danger/currency/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

func postgresCmd(ctx context.Context, conf config.Config) *cobra.Command {
	command := &cobra.Command{
		Use:   "postgres",
		Short: "run db migrations for postgres",
		RunE: func(cmd *cobra.Command, args []string) error {
			//nolint:wrapcheck
			return cmd.Usage()
		},
	}

	command.AddCommand(up(ctx, conf, postgres))
	command.AddCommand(down(ctx, conf, postgres))
	command.AddCommand(step(ctx, conf, postgres))
	command.AddCommand(version(ctx, conf, postgres))

	return command
}

func postgres(ctx context.Context, conf config.Config) (*migrate.Migrate, error) {
	b, err := build.New(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("creating builder: %w", err)
	}

	return b.PostgresMigration()
}
