package build

import (
	"github.com/3Danger/currency/internal/repo/currency/redis"
	"github.com/3Danger/currency/internal/services/fetcher/crypto_pair"
	"github.com/3Danger/currency/internal/services/fetcher/fiat"
	"github.com/3Danger/currency/pkg/cronworker"
)

func (b *Builder) NewServiceFetcherFiat() cronworker.Worker {
	var (
		repo    = redis.NewRepo(b.redis)
		httpCli = b.NewCurrencyClient()
	)

	svc := fiat.NewService(repo, httpCli)

	return cronworker.Worker{
		Execute:  svc.Process,
		Schedule: b.cnf.Workers.Updater.UpdateShedule,
		Name:     "fetcher",
	}
}

func (b *Builder) NewServiceFetcherCrypto() cronworker.Worker {
	var (
		repo    = redis.NewRepo(b.redis)
		httpCli = b.NewCurrencyClient()
	)

	svc := fiat.NewService(repo, httpCli)

	return cronworker.Worker{
		Execute:  svc.Process,
		Schedule: b.cnf.Workers.Updater.UpdateShedule,
		Name:     "fetcher",
	}
}
