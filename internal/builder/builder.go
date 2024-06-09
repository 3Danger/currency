package builder

import (
	"context"
	"fmt"
	"sync"

	"github.com/3Danger/currency/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Builder struct {
	cnf config.Config

	redis *redis.Client

	rw       sync.Mutex
	shutdown []func(ctx context.Context) error
}

func New(ctx context.Context, cnf config.Config) (*Builder, error) {
	b := Builder{cnf: cnf} //nolint:exhaustruct

	redisClient := redis.NewClient(cnf.Redis.Options())

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	b.redis = redisClient

	return &b, nil
}

func (b *Builder) addToShutdown(f func(ctx context.Context) error) {
	b.rw.Lock()
	b.shutdown = append(b.shutdown, f)
	b.rw.Unlock()
}

func (b *Builder) Shutdown(ctx context.Context) error {
	b.rw.Lock()
	defer b.rw.Unlock()

	for i := range b.shutdown {
		// в обратном порядке
		f := b.shutdown[len(b.shutdown)-1-i]

		if err := f(ctx); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("shundown")
		}
	}

	b.shutdown = b.shutdown[:0]
}
