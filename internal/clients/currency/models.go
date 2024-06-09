package currency

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Currency struct {
	Code      string
	RateToUSD decimal.Decimal
}

type Code string

const (
	CodeEUR  = Code("EUR")
	CodeUSD  = Code("USD")
	CodeCNY  = Code("CNY")
	CodeUSDT = Code("USDT")
	CodeUSDC = Code("USDC")
	CodeETH  = Code("ETH")
)

type Response struct {
	Results map[Code]decimal.Decimal `json:"results"`
	Updated Time                     `json:"updated"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	tt, err := time.Parse(time.DateTime, string(data))

	if err != nil {
		return fmt.Errorf("parsing time: %w", err)
	}

	t.Time = tt

	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(time.DateTime)), nil
}
