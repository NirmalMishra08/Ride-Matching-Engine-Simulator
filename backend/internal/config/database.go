package config

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase(ctx context.Context, config *Config) (*pgxpool.Pool, error) {

	count := 0

	for {
		conn, err := pgxpool.New(ctx, config.DB_URL)
		if err == nil {
			err := conn.Ping(ctx)
			if err == nil {
				logger.Info("connected to postgres!")
				return conn, nil
			}
		}
		count++

		logger.Info("not able to connect to postres %v", err)
		if count == 5 {
			logger.Info("retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			count = 0
		}
	}

}
