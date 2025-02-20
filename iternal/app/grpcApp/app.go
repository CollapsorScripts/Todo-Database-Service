package grpcApp

import (
	"databaseService/pkg/config"
	"databaseService/pkg/logger"
	Service "databaseService/pkg/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	cfg        *config.Config
}

func New(cfg *config.Config) *App {
	gRPCServer := grpc.NewServer()

	Service.Register(gRPCServer)

	return &App{
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(any(err))
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.GRPC.Port))
	if err != nil {
		return fmt.Errorf("ошибка при попытке прослушать порт: %w", err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("ошибка при запуске сервера: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	logger.Info("Остановка сервера")
	a.gRPCServer.GracefulStop()
	//TODO: убрать после фикса
	a.gRPCServer.Stop()
}
