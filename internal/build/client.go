package build

import (
	"net/http"

	"github.com/3Danger/currency/internal/clients/currency"
)

func (b *Builder) NewCurrencyClient() *currency.Client {
	cnf := b.cnf.Client

	httpClient := http.DefaultClient

	return currency.NewClient(httpClient, cnf.Host, cnf.Token)
}
