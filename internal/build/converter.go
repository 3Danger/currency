package build

import (
	"github.com/3Danger/currency/internal/repo/currency/redis"
	"github.com/3Danger/currency/internal/services/converter"
)

func (b *Builder) NewServiceConverter() converter.Service {
	repo := redis.NewRepo(b.redis)

	return converter.NewService(repo)
}
