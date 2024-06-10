package models

import (
	"strings"

	"github.com/3Danger/currency/pkg/time"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Currency struct {
	Code      Code
	RateToUSD decimal.Decimal
	Updated   time.Time[time.LayoutDateTime]
}

type CurrencyPair struct {
	FromCode Code
	Rate     decimal.Decimal
	ToCode   Code
	Updated  time.Time[time.LayoutDateTime]
}

type MapPossiblePairs map[Code]Code

type Pair string

func (p Pair) String() string { return string(p) }

func (p Pair) SplitCodes() (from, to Code, _ error) {
	split := strings.Split(string(p), "/")

	if len(split) != 2 {
		return "", "", errors.New("invalid pairs code")
	}

	return Code(split[0]), Code(split[1]), nil
}

func JoinCodes(a, b Code) Pair {
	return Pair(string(a) + "/" + string(b))
}
