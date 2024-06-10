package currency

import (
	"context"

	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/pkg/time"
	"github.com/shopspring/decimal"
)

type Repo interface {
	SetCurrenciesFiat(ctx context.Context, currencies []*models.Currency) error
	SetCryptoPrices(ctx context.Context, pairsRate []*models.CurrencyPair) error
	CurrencyPriceByPair(ctx context.Context, pair models.Pair) (*decimal.Decimal, *time.Time[time.LayoutDateTime], error)
	Currency(ctx context.Context, code models.Code) (*models.Currency, error)
	ListCodes(ctx context.Context) ([]models.Code, error)
}
