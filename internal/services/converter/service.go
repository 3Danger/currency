package converter

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/models"
	"github.com/shopspring/decimal"
)

type repo interface {
	CurrencyPriceByPair(ctx context.Context, pair models.Pair) (*decimal.Decimal, error)
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

	rate, err := s.r.CurrencyPriceByPair(ctx, models.JoinCodes(from, to))
	if err != nil {
		return decimal.Decimal{}, nil, fmt.Errorf("getting rate: %w", err)
	}

	var mediatorCode *models.Code
	if rate == nil {
		rate, mediatorCode, err = s.alternativePair(ctx, from, to)
		if err != nil {
			return decimal.Decimal{}, nil, fmt.Errorf("getting rate: %w", err)
		}

		if rate == nil {
			return decimal.Decimal{}, nil, models.ErrCodeNotFound
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
	)
	for _, code := range fiatCodes {
		var err error

		to = code

		rate, err = s.r.CurrencyPriceByPair(ctx, models.JoinCodes(from, to))
		if err != nil {
			return nil, nil, fmt.Errorf("getting rate: %w", err)
		}

		if rate != nil {
			break
		}
	}

	if rate == nil {
		return nil, nil, fmt.Errorf("no currency rate found")
	}

	rub, err := s.r.Currency(ctx, curFiat)
	if err != nil {
		return nil, nil, fmt.Errorf("getting currency rate: %w", err)
	}

	if rub == nil {
		return nil, nil, nil
	}

	eur, err := s.r.Currency(ctx, to)
	if err != nil {
		return nil, nil, fmt.Errorf("getting currency rate: %w", err)
	}

	if eur == nil {
		return nil, nil, nil
	}

	resultRate := (rub.RateToUSD.Div(eur.RateToUSD)).Mul(*rate)

	return &resultRate, &to, nil
}
