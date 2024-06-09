package rest

import (
	"github.com/3Danger/currency/internal/models"
	"github.com/shopspring/decimal"
)

type Result struct {
	Result       decimal.Decimal `json:"result"`
	MediatorCode *models.Code    `json:"mediator_code,omitempty"`
}
