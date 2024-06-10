package config

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Postgres struct {
	Host     string        `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     string        `envconfig:"POSTGRES_PORT" default:"5432"`
	User     string        `envconfig:"POSTGRES_USER" default:"root"`
	Password string        `envconfig:"POSTGRES_PASSWORD"`
	Database string        `envconfig:"POSTGRES_DATABASE"`
	SSL      string        `envconfig:"POSTGRES_SSL" default:"disable"`
	Timeout  time.Duration `envconfig:"POSTGRES_TIMEOUT" default:"1m"`

	MaxOpenConns    int           `envconfig:"POSTGRES_MAX_OPEN_CONNS" default:"12"`
	MaxIdleConns    int           `envconfig:"POSTGRES_MAX_IDLE_CONNS" default:"7"`
	ConnMaxLifetime time.Duration `envconfig:"POSTGRES_CONN_MAX_LIFETIME" default:"30m"`
}

func (p Postgres) DSN() string {
	dsn := strings.Builder{}
	dsn.WriteString("postgres://")

	if p.User != "" {
		dsn.WriteString(p.User)

		if p.Password != "" {
			dsn.WriteString(fmt.Sprintf(":%s", p.Password))
		}

		dsn.WriteString("@")
	}

	dsn.WriteString(fmt.Sprintf("%s/", net.JoinHostPort(p.Host, p.Port)))

	if p.Database != "" {
		dsn.WriteString(p.Database)
	}

	dsn.WriteString(fmt.Sprintf("?sslmode=%s", p.SSL))

	return dsn.String()
}
