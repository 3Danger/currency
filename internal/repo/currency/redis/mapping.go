package redis

import (
	"fmt"
	"strings"

	"github.com/3Danger/currency/pkg/time"
	"github.com/shopspring/decimal"
)

const separator = "^"

func splitRateDate(data string) (*decimal.Decimal, *time.Time[time.LayoutDateTime], error) {
	results := strings.Split(data, separator)
	if len(results) != 2 {
		return nil, nil, fmt.Errorf("invalid result: %s", data)
	}

	rate, err := decimal.NewFromString(results[0])
	if err != nil {
		return nil, nil, fmt.Errorf("converting to decimal: %w", err)
	}

	updated, err := time.NewFromString[time.LayoutDateTime](results[1])
	if err != nil {
		return nil, nil, fmt.Errorf("converting to date: %w", err)
	}

	return &rate, &updated, nil
}

func joinRateToDate(rate decimal.Decimal, updated time.Time[time.LayoutDateTime]) string {
	return rate.String() + separator + updated.UTC().String()
}
