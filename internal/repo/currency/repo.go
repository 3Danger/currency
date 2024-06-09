package currency

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type Repo interface {
	Currency(ctx context.Context, code string) (*Currency, error)
	Upsert(ctx context.Context, c []*Currency) error
}

type Currency struct {
	Code      string
	RateToUsd decimal.Decimal
	UpdatedAt time.Time
}
