package cmd

import (
	"github.com/3Danger/currency/internal/build"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func restCmd(ctx context.Context, builder *build.Builder) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run rest server",
		RunE:  runRest(ctx, builder),
	}
}

func runRest(ctx context.Context, builder *build.Builder) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		ctx, cancel := context.WithCancel(ctx)

		api := builder.ConfigureAPI(ctx)

		go func() {
			if err := api(ctx); err != nil {
				zerolog.Ctx(ctx).Err(err).Msg("api initialization failed")
			}

			cancel()
		}()
		zerolog.Ctx(ctx).Info().Str("entity", "rest").Msg("starting rest")
		<-ctx.Done()

		zerolog.Ctx(ctx).Info().Str("entity", "rest").Msg("stopping rest")

		return nil
	}
}
