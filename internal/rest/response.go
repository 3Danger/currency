package rest

import (
	"github.com/3Danger/currency/internal/models"
	"github.com/shopspring/decimal"
)

// Result представляет собой ответ на конвертацию
type Result struct {
	Result       decimal.Decimal `json:"result"                  example:"121.000"`
	MediatorCode *models.Code    `json:"mediator_code,omitempty" example:"USD"`
}

// Error представляет собой объект ошибки
type Error struct {
	Message string `json:"message" example:"произошла такая-то ошибка"`
}
