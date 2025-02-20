package database

import (
	"context"
	"databaseService/pkg/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var globalPool *pgxpool.Pool

// New - создает новый пул соединений
func New(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	cfgPgx, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	//Настройки пула
	{
		cfgPgx.MaxConns = 10
		cfgPgx.MinConns = 2
		cfgPgx.MaxConnIdleTime = 5 * time.Minute
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, cfgPgx)
	if err != nil {
		return nil, err
	}

	// Чекаем соединение
	if err = dbPool.Ping(ctx); err != nil {
		return nil, err
	}

	globalPool = dbPool

	return dbPool, nil
}

// GetPool - возвращает пул соединений с БД
func GetPool() *pgxpool.Pool {
	return globalPool
}
