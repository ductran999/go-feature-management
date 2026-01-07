package postgresql

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func New(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	address := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		address,
		cfg.DBName,
		cfg.SSLMode,
	)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = 10
	poolCfg.MinConns = 2
	poolCfg.MaxConnIdleTime = time.Minute * 5

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
