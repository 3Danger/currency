package postgres

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/internal/repo/currency"
	"github.com/3Danger/currency/internal/repo/currency/postgres/query"
	"github.com/3Danger/currency/pkg/time"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type repo struct {
	q *query.Queries
}

func NewRepo(db query.DBTX) currency.Repo {
	return &repo{
		q: query.New(db),
	}
}

func (r *repo) SetCurrenciesFiat(ctx context.Context, currencies []*models.Currency) error {
	batch := r.q.SetCurrenciesFiat(ctx, lo.Map(currencies,
		func(item *models.Currency, _ int) query.SetCurrenciesFiatParams {
			return query.SetCurrenciesFiatParams{
				Code:      query.FiatCode(item.Code),
				UpdatedAt: item.Updated.Time,
				RateToUsd: item.RateToUSD,
			}
		}))

	return process(batch)
}

func (r *repo) SetCryptoPrices(ctx context.Context, pairsRate []*models.CurrencyPair) error {
	batch := r.q.SetCryptoPrices(ctx,
		lo.Map(pairsRate, func(item *models.CurrencyPair, _ int) query.SetCryptoPricesParams {
			return query.SetCryptoPricesParams{
				CodeCrypto: query.CryptoCode(item.FromCode),
				CodeFiat:   query.FiatCode(item.ToCode),
				UpdatedAt:  item.Updated.Time,
				Rate:       item.Rate,
			}
		}))

	return process(batch)
}

func (r *repo) CurrencyPriceByPair(ctx context.Context, pair models.Pair) (*decimal.Decimal, *time.Time[time.LayoutDateTime], error) {
	crypto, fiat, err := pair.SplitCodes()
	if err != nil {
		return nil, nil, fmt.Errorf("splitting pair: %w", err)
	}

	row, err := r.q.CurrencyPriceByPair(ctx, query.CurrencyPriceByPairParams{
		CodeFiat:   query.FiatCode(fiat),
		CodeCrypto: query.CryptoCode(crypto),
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("making query: %w", err)
	}

	updatedAt := time.NewFrom[time.LayoutDateTime](row.UpdatedAt)

	return &row.Rate, &updatedAt, nil
}

func (r *repo) Currency(ctx context.Context, code models.Code) (*models.Currency, error) {
	row, err := r.q.Currency(ctx, query.FiatCode(code))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("making query: %w", err)
	}

	updatedAt := time.NewFrom[time.LayoutDateTime](row.UpdatedAt)

	return &models.Currency{
		Code:      models.Code(row.Code),
		RateToUSD: row.RateToUsd,
		Updated:   updatedAt,
	}, nil
}
