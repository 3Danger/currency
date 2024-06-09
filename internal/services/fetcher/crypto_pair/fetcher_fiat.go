package fiat

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/models"
)

type Repo interface {
	SetCurrenciesFiat(ctx context.Context, currencies []*models.Currency) error
}

type Client interface {
	CurrenciesFiat(ctx context.Context, codes []models.Code) ([]*models.Currency, error)
}

type service struct {
	repo   Repo
	client Client
}

var allowFiatCodes = []models.Code{
	models.CodeFiatUSD,
	models.CodeFiatEUR,
	models.CodeFiatCNY,
}
var allowCryptoCodes = []models.Code{
	models.CodeCryptoUSDT,
	models.CodeCryptoUSDC,
	models.CodeCryptoETH,
}

//USDC/EUR

func NewService(repo Repo, client Client) *service {
	return &service{
		repo:   repo,
		client: client,
	}
}

func (s *service) Process(ctx context.Context) error {
	currencies, err := s.client.CurrenciesFiat(ctx, allowFiatCodes)
	if err != nil {
		return fmt.Errorf("getting currencies from client: %w", err)
	}

	if err := s.repo.SetCurrenciesFiat(ctx, currencies); err != nil {
		return fmt.Errorf("setting currencies: %w", err)
	}

	return nil
}
