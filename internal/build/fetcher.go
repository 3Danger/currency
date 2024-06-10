package build

import (
	"github.com/3Danger/currency/internal/repo/currency/postgres"
	"github.com/3Danger/currency/internal/services/fetcher/crypto"
	"github.com/3Danger/currency/internal/services/fetcher/fiat"
	"github.com/3Danger/currency/pkg/cronworker"
)

func (b *Builder) NewServiceFetcherFiat() cronworker.Worker {
	var (
		//repo    = redis.NewRepo(b.redis)
		repo    = postgres.NewRepo(b.pgx)
		httpCli = b.NewCurrencyClient()
	)

	svc := fiat.NewService(repo, httpCli)

	return cronworker.Worker{
		Execute:  svc.Process,
		Schedule: b.cnf.Workers.Updater.UpdateSchedule,
		Name:     "fetcher_fiat",
	}
}

func (b *Builder) NewServiceFetcherCryptoPrices() cronworker.Worker {
	var (
		//repo    = redis.NewRepo(b.redis)
		repo    = postgres.NewRepo(b.pgx)
		httpCli = b.NewCurrencyClient()
	)

	svc := crypto.NewService(repo, httpCli)

	return cronworker.Worker{
		Execute:  svc.Process,
		Schedule: b.cnf.Workers.Updater.UpdateSchedule,
		Name:     "fetcher_crypto_prices",
	}
}
