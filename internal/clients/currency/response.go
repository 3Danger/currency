package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func processingResponse[T any](resp *http.Response) (*T, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	result := new(T)

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return result, nil
}
