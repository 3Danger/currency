package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/samber/lo"
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

func (c *Client) Currency(ctx context.Context, codes ...Code) (*Response, error) {
	req, err := makeRequestFeetchMulti(ctx, c.host, c.token, codes...)
	if err != nil {
		return nil, fmt.Errorf("making fetch multi: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}

	result, err := processingResponse[Response](resp)
	if err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return result, nil
}

func processingResponse(resp *http.Response) (*Response, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	result := new(Response)

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return result, nil
}

func makeRequestFeetchMulti(ctx context.Context, host, token string, codes ...Code) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", host+"/fetch-multi", nil)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	req.SetPathValue("api_key", token)
	req.SetPathValue("from", string(CodeUSD))
	codesString := strings.Join(lo.Map(codes, func(item Code, _ int) string { return string(item) }), ",")
	req.SetPathValue("to", codesString)

	return req, nil
}
