package database

import (
	"context"
	"fmt"

	"github.com/mixdone/fly-api/internal/config"

	"github.com/jackc/pgx/v5"
)

func ConnectToDB(cfg *config.Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User,
		"2388",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}
