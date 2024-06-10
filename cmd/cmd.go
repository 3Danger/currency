package cmd

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/build"
	"github.com/3Danger/currency/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Run(ctx context.Context, conf config.Config) error {
	root := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	builder, err := build.New(ctx, conf)
	if err != nil {
		return fmt.Errorf("creating builder: %w", err)
	}

	root.AddCommand(
		restCmd(ctx, builder),
		workersCmd(ctx, builder),
		migrateCmd(ctx, builder.Config()),
	)

	return errors.Wrap(root.ExecuteContext(ctx), "run application")
}
