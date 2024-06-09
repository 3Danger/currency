package currency

import (
	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/pkg/time"
	"github.com/shopspring/decimal"
)

type ResponseFiat struct {
	Results map[string]decimal.Decimal     `json:"results"`
	Updated time.Time[time.LayoutDateTime] `json:"updated"`
}

type ResponsePrices struct {
	Prices map[models.Pair]decimal.Decimal `json:"prices"`
}

type ResponsePossiblePairs struct {
	Pairs map[models.Pair]any `json:"pairs"`
}
