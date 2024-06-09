package currency

import (
	"context"

	"github.com/3Danger/currency/internal/models"
)

type Repo interface {
	SetCurrenciesFiat(ctx context.Context, currencies []*models.Currency) error
	Currency(ctx context.Context, code models.Code) (*models.Currency, error)
	ListCodes(ctx context.Context) ([]models.Code, error)
}
