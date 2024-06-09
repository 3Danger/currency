package models

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Currency struct {
	Code      Code
	RateToUSD decimal.Decimal
}

type CurrencyPair struct {
	FromCode Code
	Rate     decimal.Decimal
	ToCode   Code
}

type Code string

const (
	CodeFiatUSD = Code("USD")
	CodeFiatEUR = Code("EUR")
	CodeFiatCNY = Code("CNY")
)

const (
	CodeCryptoUSDT = Code("USDT")
	CodeCryptoUSDC = Code("USDC")
	CodeCryptoETH  = Code("ETH")
)

type MapPossiblePairs map[Code]Code

type Pair string

func (p Pair) SplitCodes() (from, to Code, _ error) {
	split := strings.Split(string(p), "/")

	if len(split) != 2 {
		return "", "", errors.New("invalid pairs code")
	}

	return Code(split[0]), Code(split[1]), nil
}

func joinCodes(A, B Code) Pair {
	return Pair(string(A) + "/" + string(B))
}
