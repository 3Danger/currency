package cmd

import (
	"context"

	"github.com/3Danger/currency/internal/build"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func restCmd(ctx context.Context, builder *build.Builder) *cobra.Command {
	//nolint:exhaustruct,goconst
	cmd := &cobra.Command{
		Use:   "rest",
		Short: "rest service",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			zerolog.Ctx(ctx).Info().Msg("start " + cmd.Short)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			zerolog.Ctx(ctx).Info().Msg("stop " + cmd.Short)
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			//TODO реализовать апи

			return cmd.Usage()
		},
	}

	cmd.AddCommand(
		workersFetcherFiatCmd(ctx, builder),
	)

	return cmd
}
