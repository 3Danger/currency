package build

import (
	"github.com/3Danger/currency/internal/repo/currency/postgres"
	"github.com/3Danger/currency/internal/services/converter"
)

func (b *Builder) NewServiceConverter() converter.Service {
	//repo := redis.NewRepo(b.redis)
	repo := postgres.NewRepo(b.pgx)

	return converter.NewService(repo)
}
