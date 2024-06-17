package redis

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/internal/repo/currency"
	"github.com/3Danger/currency/pkg/time"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type repo struct {
	cli *redis.Client
}

const (
	fiatCurrency   = "currencies:fiat"
	cryptoCurrency = "currencies:crypto_price"
)

func NewRepo(cli *redis.Client) currency.Repo {
	return &repo{cli: cli}
}

const ts = "_ts"

func (r *repo) SetCurrenciesFiat(ctx context.Context, currencies []*models.Currency) error {
	currencyMap := make(map[string]string, len(currencies))

	for _, cur := range currencies {
		currencyMap[cur.Code.String()] = cur.RateToUSD.String()
		currencyMap[cur.Code.String()+ts] = cur.Updated.String()
	}

	if err := r.cli.HSet(ctx, fiatCurrency, currencyMap).Err(); err != nil {
		return fmt.Errorf("hsetting to redis: %w", err)
	}

	return nil
}

func (r *repo) SetCryptoPrices(ctx context.Context, pairsRate []*models.CurrencyPair) error {
	pairsRateMap := make(map[string]string, len(pairsRate))

	for _, pair := range pairsRate {
		pairString := models.JoinCodes(pair.FromCode, pair.ToCode).String()

		pairsRateMap[pairString] = pair.Rate.String()
		pairsRateMap[pairString+ts] = pair.Updated.String()
	}

	if err := r.cli.HSet(ctx, cryptoCurrency, pairsRateMap).Err(); err != nil {
		return fmt.Errorf("hsetting to redis: %w", err)
	}

	return nil
}

func (r *repo) CurrencyPriceByPair(ctx context.Context, pair models.Pair) (
	*decimal.Decimal, *time.Time[time.LayoutDateTime], error,
) {
	row, err := r.cli.HGet(ctx, cryptoCurrency, pair.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	rowTs, err := r.cli.HGet(ctx, cryptoCurrency, pair.String()+ts).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	return splitRateDate(row, rowTs)
}

func (r *repo) Currency(ctx context.Context, code models.Code) (*models.Currency, error) {
	row, err := r.cli.HGet(ctx, fiatCurrency, string(code)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	rowTs, err := r.cli.HGet(ctx, fiatCurrency, string(code)+ts).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, fmt.Errorf("hgetting from redis: %w", err)
	}

	rateToUSD, updated, err := splitRateDate(row, rowTs)
	if err != nil {
		return nil, fmt.Errorf("parsing rate date: %w", err)
	}

	return &models.Currency{Code: code, RateToUSD: *rateToUSD, Updated: *updated}, nil
}
