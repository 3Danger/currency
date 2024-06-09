package redis

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/repo/currency"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type repo struct {
	cli *redis.Client
}

const entity = "currencies"

func NewRepo(cli *redis.Client) *repo {
	return &repo{cli: cli}
}

func (r *repo) SetCurrency(ctx context.Context, currencies []*currency.Currency) error {
	currencyMap := lo.SliceToMap(currencies,
		func(item *currency.Currency) (string, decimal.Decimal) {
			return item.Code, item.RateToUsd
		})

	if err := r.cli.HSet(ctx, entity, currencyMap).Err(); err != nil {
		return fmt.Errorf("hsetting to redis: %w", err)
	}

	return nil
}

func (r *repo) Currency(ctx context.Context, code string) (*currency.Currency, error) {
	result, err := r.cli.HGet(ctx, entity, code).Result()
	if err != nil {
		return nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	rateToUsd, err := decimal.NewFromString(result)
	if err != nil {
		return nil, fmt.Errorf("converting to decimal: %w", err)
	}

	return &currency.Currency{Code: code, RateToUsd: rateToUsd}, nil
}

func (r *repo) ListCodes(ctx context.Context) ([]string, error) {
	result, err := r.cli.Keys(ctx, entity).Result()
	if err != nil {
		return nil, fmt.Errorf("keys from redis: %w", err)
	}

	return result, nil
}
