package currency

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/3Danger/currency/internal/models"
	"github.com/samber/lo"
)

const (
	pathFiatFetchMulti    = "/fetch-multi"
	pathCryptoPairs       = "/crypto/pairs"
	pathCryptoFetchPrices = "/crypto/fetch-prices"
)

func makeRequestFiatFeetchMulti(ctx context.Context, host, token string, codes ...models.Code) (*http.Request, error) {
	paramCodes := lo.Map(codes,
		func(item models.Code, _ int) string {
			return string(item)
		},
	)

	path, err := url.JoinPath(host, pathFiatFetchMulti)
	if err != nil {
		return nil, fmt.Errorf("making path: %w", err)
	}

	pathParams := "api_key=" + token + "&from=" + string(models.CodeFiatUSD) + "&to=" + strings.Join(paramCodes, ",")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path+"?"+pathParams, nil)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	return req, nil
}

func makeRequestCryptoPossiblePairs(ctx context.Context, host, token string) (*http.Request, error) {
	path, err := url.JoinPath(host, pathCryptoPairs)
	if err != nil {
		return nil, fmt.Errorf("making path: %w", err)
	}

	pathParams := "api_key=" + token

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path+"?"+pathParams, nil)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	return req, nil
}

func makeRequestCryptoFetchPrices(ctx context.Context, host, token string, pairs []*models.Pair) (*http.Request, error) {
	path, err := url.JoinPath(host, pathCryptoFetchPrices)
	if err != nil {
		return nil, fmt.Errorf("making path: %w", err)
	}

	sbPairs := new(strings.Builder)

	count := len(pairs)
	for _, pair := range pairs {
		sbPairs.WriteString(pair.String())

		if count--; count > 0 {
			sbPairs.WriteString(",")
		}
	}

	pathParams := "api_key=" + token + "&pairs=" + sbPairs.String()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path+"?"+pathParams, nil)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	return req, nil
}
