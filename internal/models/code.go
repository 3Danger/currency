package models

import "github.com/samber/lo"

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

func (c Code) String() string { return string(c) }

func (c Code) IsValid() bool {
	return c.IsFiat() || c.IsCrypto()
}

func (c Code) IsFiat() bool {
	_, ok := fiatMap[c]

	return ok
}

func (c Code) IsCrypto() bool {
	_, ok := cryptoMap[c]

	return ok
}

var fiatMap = map[Code]struct{}{
	CodeFiatUSD: {},
	CodeFiatEUR: {},
	CodeFiatCNY: {},
}

func FiatCodes() []Code {
	return lo.MapToSlice(fiatMap, func(key Code, _ struct{}) Code {
		return key
	})
}

var cryptoMap = map[Code]struct{}{
	CodeCryptoUSDT: {},
	CodeCryptoUSDC: {},
	CodeCryptoETH:  {},
}
