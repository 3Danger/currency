package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Config struct {
	App      App
	Rest     Rest
	Redis    Redis
	Postgres Postgres
	Client   Client
	Workers  Workers
}

type App struct {
	LogLevel string `envconfig:"APP_LOG_LEVEL" default:"info"`
}

type Client struct {
	Host  string `envconfig:"HTTP_HOST"  required:"true"`
	Token string `envconfig:"HTTP_TOKEN" required:"true"`
}

type Rest struct {
	Port string `envconfig:"REST_PORT" default:":8080"`
}

type Workers struct {
	Updater struct {
		Host           string `envconfig:"WORKERS_UPDATER_HOST"            required:"true"`
		UpdateSchedule string `envconfig:"WORKERS_UPDATER_UPDATE_SCHEDULE" default:"* * * * *"`
	}
}

type Redis struct {
	Addr            string        `envconfig:"REDIS_HOST"                default:"localhost:6379"`
	DialTimeout     time.Duration `envconfig:"REDIS_DIAL_TIMEOUT"        default:"1m"`
	ReadTimeout     time.Duration `envconfig:"REDIS_READ_TIMEOUT"        default:"1m"`
	WriteTimeout    time.Duration `envconfig:"REDIS_WRITE_TIMEOUT"       default:"1m"`
	MinIdleConns    int           `envconfig:"REDIS_MIN_IDLE_CONNS"      default:"5"`
	MaxIdleConns    int           `envconfig:"REDIS_MAX_IDLE_CONNS"      default:"20"`
	MaxActiveConns  int           `envconfig:"REDIS_MAX_ACTIVE_CONNS"    default:"30"`
	ConnMaxIdleTime time.Duration `envconfig:"REDIS_CONN_MAX_IDLE_TIME"  default:"30m"`
	ConnMaxLifetime time.Duration `envconfig:"REDIS_CONN_MAX_LIFETIME"   default:"30m"`
}

func (r Redis) Options() *redis.Options {
	return &redis.Options{
		Addr:            r.Addr,
		DialTimeout:     r.DialTimeout,
		ReadTimeout:     r.ReadTimeout,
		WriteTimeout:    r.WriteTimeout,
		MinIdleConns:    r.MinIdleConns,
		MaxIdleConns:    r.MaxIdleConns,
		MaxActiveConns:  r.MaxActiveConns,
		ConnMaxIdleTime: r.ConnMaxIdleTime,
		ConnMaxLifetime: r.ConnMaxLifetime,
	}
}

func (a App) ParseLevel() (zerolog.Level, error) {
	level, err := zerolog.ParseLevel(a.LogLevel)
	if err != nil {
		return level, fmt.Errorf("parsing level: %w", err)
	}

	return level, nil
}

func Load() (Config, error) {
	cnf := Config{}

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return cnf, fmt.Errorf("loading .env: %w", err)
	}

	if err := envconfig.Process("", &cnf); err != nil {
		return cnf, fmt.Errorf("parsing env to config: %w", err)
	}

	return cnf, nil
}
