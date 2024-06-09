package crypto

import (
	"context"
	"fmt"

	"github.com/3Danger/currency/internal/models"
)

type Repo interface {
	SetCryptoPrices(ctx context.Context, pairsRate []*models.CurrencyPair) error
}

type Client interface {
	CryptoPrices(ctx context.Context, pairs []*models.Pair) ([]*models.CurrencyPair, error)
}

type service struct {
	repo   Repo
	client Client
}

func NewService(repo Repo, client Client) *service {
	return &service{
		repo:   repo,
		client: client,
	}
}

func (s *service) Process(ctx context.Context) error {
	cryptoPrices, err := s.client.CryptoPrices(ctx, allowPairs)
	if err != nil {
		return fmt.Errorf("getting cryptoPrices from client: %w", err)
	}

	if err := s.repo.SetCryptoPrices(ctx, cryptoPrices); err != nil {
		return fmt.Errorf("setting cryptoPrices: %w", err)
	}

	return nil
}
