package crypto

import "github.com/3Danger/currency/internal/models"

var allowPairs []*models.Pair

func init() {
	allowFiatCodes := [...]models.Code{
		models.CodeFiatUSD,
		models.CodeFiatEUR,
		models.CodeFiatCNY,
	}

	allowCryptoCodes := [...]models.Code{
		models.CodeCryptoUSDT,
		models.CodeCryptoUSDC,
		models.CodeCryptoETH,
	}

	// Максимально допустимое кол-во 10 пар в запросе
	allowPairs = make([]*models.Pair, 0, len(allowFiatCodes)*len(allowCryptoCodes))

	for _, crypto := range allowCryptoCodes {
		for _, fiat := range allowFiatCodes {
			pair := models.JoinCodes(crypto, fiat)

			allowPairs = append(allowPairs, &pair)
		}
	}
}
