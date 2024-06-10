// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package query

import (
	"context"
)

const currency = `-- name: Currency :one
SELECT code, updated_at, rate_to_usd FROM currency_fiat WHERE code = $1
`

func (q *Queries) Currency(ctx context.Context, code FiatCode) (CurrencyFiat, error) {
	row := q.db.QueryRow(ctx, currency, code)
	var i CurrencyFiat
	err := row.Scan(&i.Code, &i.UpdatedAt, &i.RateToUsd)
	return i, err
}

const currencyPriceByPair = `-- name: CurrencyPriceByPair :one
SELECT code_crypto, code_fiat, updated_at, rate
FROM currency_crypto_pair
WHERE code_fiat   = $1
  AND code_crypto = $2
`

type CurrencyPriceByPairParams struct {
	CodeFiat   FiatCode
	CodeCrypto CryptoCode
}

func (q *Queries) CurrencyPriceByPair(ctx context.Context, arg CurrencyPriceByPairParams) (CurrencyCryptoPair, error) {
	row := q.db.QueryRow(ctx, currencyPriceByPair, arg.CodeFiat, arg.CodeCrypto)
	var i CurrencyCryptoPair
	err := row.Scan(
		&i.CodeCrypto,
		&i.CodeFiat,
		&i.UpdatedAt,
		&i.Rate,
	)
	return i, err
}
