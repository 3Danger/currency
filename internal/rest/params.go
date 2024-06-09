package rest

import (
	"github.com/3Danger/currency/internal/models"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type BodyParams struct {
	From  *models.Code     `json:"from"  required:"true"`
	To    *models.Code     `json:"to"    required:"true"`
	Value *decimal.Decimal `json:"value" required:"true"`
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
