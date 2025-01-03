package postgres

import (
	"context"
	"fmt"
	"rocket-web/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Storage struct {
	conn *pgxpool.Pool
}

func New(ctx context.Context, log *zap.SugaredLogger, cfg *config.PostgresConfig) (*Storage, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	log.Debug("Connecting to database", zap.String("dsn", dsn))
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return &Storage{
		conn: conn,
	}, nil
}

func (s *Storage) Close() {
	s.conn.Close()
}

func (s *Storage) Ping(ctx context.Context) error {
	return s.conn.Ping(ctx)
}
