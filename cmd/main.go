package main

import (
	"fmt"
	"os"

	"github.com/3Danger/currency/internal/config"
	zlog "github.com/rs/zerolog"
)

//TODO поддерживать EUR, USD, CNY, USDT, USDC, ETH

func main() {
	if err := app(); err != nil {
		panic(err)
	}
}

func app() error {
	cnf, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	zlog.New(os.Stdout).
		Level(logLevel).
		With().Timestamp().Stack().Caller().
		Logger()

}
