package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/3Danger/currency/cmd"
	"github.com/3Danger/currency/internal/config"
	zlog "github.com/rs/zerolog"
)

// @title Сервис конвертации валюты
// @version 1.0.0
// @description API сервиса Currency

// @host localhost:8080
// @BasePath /api
func main() {
	if err := app(); err != nil {
		fmt.Println(err)
	}
}

func app() error {
	time.Local = time.UTC

	cnf, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logLevel, err := cnf.App.ParseLevel()
	if err != nil {
		zlog.DefaultContextLogger.Err(err).Msg("parse level")

		logLevel = zlog.InfoLevel
	}

	ctx := zlog.New(os.Stdout).
		Level(logLevel).
		With().Timestamp().Stack().Caller().
		Logger().WithContext(context.Background())

	if err := cmd.Run(ctx, cnf); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}
