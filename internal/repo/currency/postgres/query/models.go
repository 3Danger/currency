// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package query

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type Currency struct {
	Code      string
	RateToUsd decimal.Decimal
	UpdatedAt pgtype.Timestamp
}