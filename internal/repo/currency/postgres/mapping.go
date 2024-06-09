package postgres

import (
	"github.com/3Danger/currency/internal/repo/currency"
	"github.com/3Danger/currency/internal/repo/currency/postgres/query"
)

func mapCurrencyQueryToRepo(cur *query.Currency) *currency.Currency {
	return &currency.Currency{
		Code:      cur.Code,
		RateToUsd: cur.RateToUsd,
		UpdatedAt: cur.UpdatedAt,
	}
}

func mapCurrencyRepoToQuery(cur *currency.Currency) *query.Currency {
	return &query.Currency{
		Code:      cur.Code,
		RateToUsd: cur.RateToUsd,
		UpdatedAt: cur.UpdatedAt,
	}
}
