package main

import (
	"context"
	"databaseService/iternal/app"
	"databaseService/pkg/config"
	"databaseService/pkg/database"
	"databaseService/pkg/logger"
	"databaseService/pkg/utilities"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Инициализация конфигурации
	cfg := config.MustLoad()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if err := logger.New(cfg); err != nil {
		fmt.Printf("Ошибка при инициализации логера: %v\n", err)
	}

	logger.Info("Конфигурация: \n%s", utilities.ToJSON(cfg))

	//Миграции
	if cfg.Database.Migrations {
		err := database.MakeMigrations(cfg)
		if err != nil {
			logger.Error("Ошибка создания таблицы: %v", err)
		}
	}

	//Создаем пул соединений с БД
	dbPool, err := database.New(cfg)
	if err != nil {
		panic(any(fmt.Sprintf("Ошибка при инициализации подключения к БД: %v", err)))
	}

	//Инициализация приложения
	application := app.New(cfg, dbPool)

	//Правильное завершение сервиса
	{
		wait := time.Second * 15

		// Запуск сервера в отдельном потоке
		go application.GRPC.MustRun()

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

		<-c

		_, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		defer dbPool.Close()
		application.GRPC.Stop()
		os.Exit(0)
	}
}
