package rest

import (
	"github.com/3Danger/currency/internal/models"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// BodyParams представляет собой запрос на конвертацию
type BodyParams struct {
	From  *models.Code     `json:"from"  example:"USD"    required:"true"`
	To    *models.Code     `json:"to"    example:"USDT"   required:"true"`
	Value *decimal.Decimal `json:"value" example:"20.000" required:"true"`
}

func (p *BodyParams) Validate() error {
	if p.From == nil {
		return errors.New("from is required")
	}

	if p.To == nil {
		return errors.New("to is required")
	}

	if !p.From.IsValid() || !p.To.IsValid() {
		return models.ErrCodeInvalid
	}

	if p.Value == nil {
		return errors.New("value is required")
	}

	return nil
}
