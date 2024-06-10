package build

import (
	"context"
	"fmt"
	"sync"

	"github.com/3Danger/currency/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Builder struct {
	cnf config.Config

	//redis *redis.Client
	pgx *pgxpool.Pool

	rw       sync.Mutex
	shutdown []func(ctx context.Context) error
}

func New(ctx context.Context, cnf config.Config) (*Builder, error) {
	b := Builder{cnf: cnf}

	//redisClient := redis.NewClient(cnf.Redis.Options())

	//if err := redisClient.Ping(ctx).Err(); err != nil {
	//	return nil, fmt.Errorf("ping redis: %w", err)
	//}

	//b.redis = redisClient

	pool, err := pgxpool.New(ctx, b.cnf.Postgres.DSN())
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	b.pgx = pool

	return &b, nil
}

func (b *Builder) Config() config.Config { return b.cnf }

func (b *Builder) addToShutdown(f func(ctx context.Context) error) {
	b.rw.Lock()
	b.shutdown = append(b.shutdown, f)
	b.rw.Unlock()
}

func (b *Builder) Shutdown(ctx context.Context) {
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
