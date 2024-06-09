package cmd

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/build"
	"github.com/3Danger/currency/pkg/cronworker"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func workersCmd(ctx context.Context, builder *build.Builder) *cobra.Command {
	//nolint:exhaustruct,goconst
	cmd := &cobra.Command{
		Use: "fetchers",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			zerolog.Ctx(ctx).Info().Msg("start " + cmd.Short)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			zerolog.Ctx(ctx).Info().Msg("stop " + cmd.Short)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage() //nolint:wrapcheck
		},
	}

	cmd.AddCommand(
		workersFetcherAllCmd(ctx, builder),
		workersFetcherFiatCmd(ctx, builder),
		workersFetcherCryptoCmd(ctx, builder),
	)

	return cmd
}

func workersFetcherAllCmd(ctx context.Context, builder *build.Builder) *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "all",
		Short: "",
		RunE: runWorkers(ctx, []cronworker.Worker{
			builder.NewServiceFetcherFiat(),
			builder.NewServiceFetcherCryptoPrices(),
		}),
	}
}

func workersFetcherFiatCmd(ctx context.Context, builder *build.Builder) *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "fiat",
		Short: "",
		RunE: runWorkers(ctx, []cronworker.Worker{
			builder.NewServiceFetcherFiat(),
		}),
	}
}

func workersFetcherCryptoCmd(ctx context.Context, builder *build.Builder) *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "crypto",
		Short: "",
		RunE: runWorkers(ctx, []cronworker.Worker{
			builder.NewServiceFetcherCryptoPrices(),
		}),
	}
}

func runWorkers(ctx context.Context, worker []cronworker.Worker) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		tm := cronworker.NewTaskManager()

		for _, w := range worker {
			zerolog.Ctx(ctx).Info().Str("entity", "cron-task-manager").Str("name", w.Name).Msg("add task")

			if err := tm.AddWorker(ctx, w); err != nil {
				return fmt.Errorf("adding cron-worker: %w", err)
			}
		}

		zerolog.Ctx(ctx).Info().Str("entity", "cron-task-manager").Msg("starting tasks")
		tm.Start()
		<-ctx.Done()

		zerolog.Ctx(ctx).Info().Str("entity", "cron-task-manager").Msg("stopping tasks")
		tm.Stop()

		return nil
	}
}
