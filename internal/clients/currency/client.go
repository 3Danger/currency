package currency

import (
	"context"
	"fmt"
	"net/http"

	"github.com/3Danger/currency/internal/models"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient doer
	token      string
	host       string
}

func NewClient(httpClient doer, host, token string) *Client {
	return &Client{
		token:      token,
		host:       host,
		httpClient: httpClient,
	}
}

func (c *Client) CurrenciesFiat(ctx context.Context, codes []models.Code) ([]*models.Currency, error) {
	req, err := makeRequestFiatFetchMulti(ctx, c.host, c.token, codes...)
	if err != nil {
		return nil, fmt.Errorf("making fetch multi: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}

	result, err := processingResponse[ResponseFiat](resp)
	if err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return lo.MapToSlice(result.Results, func(key string, value decimal.Decimal) *models.Currency {
		return &models.Currency{
			Code:      models.Code(key),
			RateToUSD: value,
		}
	}), nil
}

func (c *Client) PossiblePairs(ctx context.Context) (models.MapPossiblePairs, error) {
	req, err := makeRequestCryptoPossiblePairs(ctx, c.host, c.token)
	if err != nil {
		return nil, fmt.Errorf("making fetch multi: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}

	result, err := processingResponse[ResponsePossiblePairs](resp)
	if err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	possiblePairs := make(models.MapPossiblePairs, len(result.Pairs))

	for pair := range result.Pairs {
		crypto, fiat, err := pair.SplitCodes()
		if err != nil {
			return nil, fmt.Errorf("parsing pair: %w", err)
		}

		possiblePairs[crypto] = fiat
	}

	return possiblePairs, nil
}

func (c *Client) CryptoPrices(ctx context.Context, pairs []*models.Pair) ([]*models.CurrencyPair, error) {
	if len(pairs) > 10 {
		return nil, fmt.Errorf("too many pairs provided")
	}

	req, err := makeRequestCryptoFetchPrices(ctx, c.host, c.token, pairs)
	if err != nil {
		return nil, fmt.Errorf("making fetch multi: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}

	result, err := processingResponse[ResponsePrices](resp)
	if err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	cyrrencyPairs := make([]*models.CurrencyPair, 0, len(result.Prices))

	for pair, value := range result.Prices {
		from, to, err := pair.SplitCodes()
		if err != nil {
			return nil, fmt.Errorf("splitting pair: %w", err)
		}

		cyrrencyPairs = append(cyrrencyPairs, &models.CurrencyPair{
			FromCode: from,
			Rate:     value,
			ToCode:   to,
		})
	}

	return cyrrencyPairs, nil
}
