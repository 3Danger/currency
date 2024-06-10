package converter

import (
	"context"
	"fmt"
	t "time"

	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/pkg/time"
	"github.com/shopspring/decimal"
)

type repo interface {
	CurrencyPriceByPair(ctx context.Context, pair models.Pair) (
		*decimal.Decimal, *time.Time[time.LayoutDateTime], error,
	)
	Currency(ctx context.Context, code models.Code) (*models.Currency, error)
}

type Service interface {
	Convert(ctx context.Context, pair models.Pair, value decimal.Decimal) (
		result decimal.Decimal, mediatorCode *models.Code, err error,
	)
}

type service struct {
	r repo
}

func NewService(r repo) Service {
	return &service{
		r: r,
	}
}

func (s *service) Convert(
	ctx context.Context, pair models.Pair, value decimal.Decimal,
) (decimal.Decimal, *models.Code, error) {
	from, to, err := pair.SplitCodes()
	if err != nil {
		return decimal.Decimal{}, nil, fmt.Errorf("parsing pair code: %w", err)
	}

	if from.IsFiat() == to.IsFiat() {
		if from.IsFiat() {
			return decimal.Decimal{}, nil, models.ErrFiatToFiatConvertForbidden
		}

		return decimal.Decimal{}, nil, models.ErrCryptoToCryptoConvertForbidden
	}

	isNeedSwap := from.IsFiat() && to.IsCrypto()
	if isNeedSwap {
		from, to = to, from
	}

	rate, updated, err := s.r.CurrencyPriceByPair(ctx, models.JoinCodes(from, to))
	if err != nil {
		return decimal.Decimal{}, nil, fmt.Errorf("getting rate: %w", err)
	}

	if updated != nil && t.Since(updated.Time) > t.Minute {
		return decimal.Decimal{}, nil, models.ErrCurrencyIsDeprecated
	}

	var mediatorCode *models.Code
	if rate == nil {
		rate, mediatorCode, err = s.alternativePair(ctx, from, to)
		if err != nil {
			return decimal.Decimal{}, nil, fmt.Errorf("getting rate: %w", err)
		}

		if rate == nil {
			return decimal.Decimal{}, nil, models.ErrCurrencyNotFound
		}
	}

	var result decimal.Decimal

	if isNeedSwap {
		result = value.Div(*rate)
	} else {
		result = value.Mul(*rate)
	}

	return result, mediatorCode, nil
}

var fiatCodes = models.FiatCodes()

func (s *service) alternativePair(ctx context.Context, from, to models.Code) (*decimal.Decimal, *models.Code, error) {
	var (
		curFiat = to
		rate    *decimal.Decimal
		updated *time.Time[time.LayoutDateTime]
	)
	for _, code := range fiatCodes {
		var err error

		to = code

		rate, updated, err = s.r.CurrencyPriceByPair(ctx, models.JoinCodes(from, to))
		if err != nil {
			return nil, nil, fmt.Errorf("getting rate: %w", err)
		}

		if rate != nil && updated != nil && t.Since(updated.Time) <= t.Minute {
			break
		}
	}

	if rate == nil || updated == nil {
		return nil, nil, models.ErrCurrencyNotFound
	}

	if t.Since(updated.Time) > t.Minute {
		return nil, nil, models.ErrCurrencyIsDeprecated
	}

	rub, err := s.r.Currency(ctx, curFiat)
	if err != nil {
		return nil, nil, fmt.Errorf("getting currency rate: %w", err)
	}

	if rub == nil {
		return nil, nil, models.ErrCurrencyNotFound
	}

	eur, err := s.r.Currency(ctx, to)
	if err != nil {
		return nil, nil, fmt.Errorf("getting currency rate: %w", err)
	}

	if eur == nil {
		return nil, nil, models.ErrCurrencyNotFound
	}

	resultRate := (rub.RateToUSD.Div(eur.RateToUSD)).Mul(*rate)

	return &resultRate, &to, nil
}
