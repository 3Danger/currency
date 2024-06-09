package redis

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/internal/repo/currency"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type repo struct {
	cli *redis.Client
}

const (
	fiatCurrency   = "currencies:fiat"
	cryptoCurrency = "currencies:crypto_price"
	possiblePairs  = "currencies:possible_pairs"
)

func NewRepo(cli *redis.Client) currency.Repo {
	return &repo{cli: cli}
}

func (r *repo) SetCurrenciesFiat(ctx context.Context, currencies []*models.Currency) error {
	currencyMap := lo.SliceToMap(currencies,
		func(item *models.Currency) (string, string) {
			return item.Code.String(), item.RateToUSD.String()
		})

	if err := r.cli.HSet(ctx, fiatCurrency, currencyMap).Err(); err != nil {
		return fmt.Errorf("hsetting to redis: %w", err)
	}

	return nil
}

func (r *repo) SetCryptoPrices(ctx context.Context, pairsRate []*models.CurrencyPair) error {
	pairsRateMap := lo.SliceToMap(pairsRate,
		func(item *models.CurrencyPair) (string, string) {
			return models.JoinCodes(item.FromCode, item.ToCode).String(), item.Rate.String()
		},
	)

	if err := r.cli.HSet(ctx, cryptoCurrency, pairsRateMap).Err(); err != nil {
		return fmt.Errorf("hsetting to redis: %w", err)
	}

	return nil
}

func (r *repo) CurrencyPriceByPair(ctx context.Context, pair models.Pair) (*decimal.Decimal, error) {
	result, err := r.cli.HGet(ctx, cryptoCurrency, pair.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	rate, err := decimal.NewFromString(result)
	if err != nil {
		return nil, fmt.Errorf("converting to decimal: %w", err)
	}

	return &rate, nil
}

func (r *repo) Currency(ctx context.Context, code models.Code) (*models.Currency, error) {
	result, err := r.cli.HGet(ctx, fiatCurrency, string(code)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	rateToUsd, err := decimal.NewFromString(result)
	if err != nil {
		return nil, fmt.Errorf("converting to decimal: %w", err)
	}

	return &models.Currency{Code: code, RateToUSD: rateToUsd}, nil
}

func (r *repo) ListCodes(ctx context.Context) ([]models.Code, error) {
	result, err := r.cli.Keys(ctx, fiatCurrency).Result()
	if err != nil {
		return nil, fmt.Errorf("keys from redis: %w", err)
	}

	return lo.Map(result, func(item string, _ int) models.Code {
		return models.Code(item)
	}), nil
}
