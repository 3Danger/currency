package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/3Danger/currency/internal/repo/currency"
	"github.com/3Danger/currency/internal/repo/currency/postgres/query"
	"github.com/jackc/pgx/v5"
)

type repo struct {
	q *query.Queries
}

func NewRepo(db *pgx.Conn) currency.Repo {
	return &repo{
		q: query.New(db),
	}
}

func (r *repo) Currency(ctx context.Context, code string) (*currency.Currency, error) {
	row, err := r.q.Currency(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("making query: %w", err)
	}

	return mapCurrencyQueryToRepo(&row), nil
}

func (r *repo) Upsert(ctx context.Context, c *currency.Currency) error {
	if err := r.q.Upsert(ctx, query.UpsertParams{
		Code:      c.Code,
		Ratetousd: c.RateToUsd,
	}); err != nil {
		return fmt.Errorf("making query: %w", err)
	}
}
