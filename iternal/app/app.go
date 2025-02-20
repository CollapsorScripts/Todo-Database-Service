package app

import (
	"databaseService/iternal/app/grpcApp"
	"databaseService/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	GRPC *grpcApp.App
	cfg  *config.Config
	db   *pgxpool.Pool
}

func New(cfg *config.Config, db *pgxpool.Pool) *App {
	grpcApp := grpcApp.New(cfg)

	return &App{GRPC: grpcApp, db: db}
}
