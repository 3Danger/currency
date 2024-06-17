package redis

import (
	"fmt"

	"github.com/3Danger/currency/pkg/time"
	"github.com/shopspring/decimal"
)

func splitRateDate(rate, rowTs string) (
	*decimal.Decimal, *time.Time[time.LayoutDateTime], error,
) {
	dec, err := decimal.NewFromString(rate)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding rate: %w", err)
	}

	timestamp, err := time.NewFromString[time.LayoutDateTime](rowTs)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding timestamp: %w", err)
	}

	return &dec, &timestamp, nil
}
