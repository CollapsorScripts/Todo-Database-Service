package database

import (
	"context"
	"databaseService/pkg/config"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func createDatabase(ctx context.Context, cfg *config.Config) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close(ctx)
	}()

	// Проверяем наличие БД
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1);"
	err = conn.QueryRow(ctx, query, cfg.Database.Name).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	// Создаем БД, если не существует
	_, err = conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s;", cfg.Database.Name))
	if err != nil {
		log.Fatalf("Ошибка создания БД: %v", err)
	}

	return nil
}

// MakeMigrations - применяет миграции
func MakeMigrations(cfg *config.Config) error {
	ctx := context.Background()
	if err := createDatabase(ctx, cfg); err != nil {
		return err
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close(ctx)
	}()

	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
        created_at TIMESTAMP DEFAULT now(),
        updated_at TIMESTAMP DEFAULT now()
    );`

	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
