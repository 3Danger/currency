package models

import "github.com/shopspring/decimal"

type Currency struct {
	Code      string
	RateToUSD decimal.Decimal
}
