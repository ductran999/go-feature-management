package bootstrap

import (
	"context"
	httpAdapter "feature-flag-poc/internal/adapter/http"
	adapterPostgres "feature-flag-poc/internal/adapter/postgresql"
	"feature-flag-poc/internal/application/usecase"
	"feature-flag-poc/internal/config"
	"feature-flag-poc/internal/db/generated"
	"feature-flag-poc/internal/infra/db/postgresql"
	httpserver "feature-flag-poc/internal/server/http"
	"log"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalln("failed to load config", err)
	}

	//  Init DB
	pool, err := postgresql.New(ctx, postgresql.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Pass,
		DBName:   cfg.DB.DatName,
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalln("failed to connect db", err)
	}

	queries := generated.New(pool)
	repo := adapterPostgres.NewTodoRepository(queries)

	// usecase
	uc := usecase.NewListTodoUsecase(repo)

	// http adapter
	handler := httpAdapter.NewTodoHandler(uc)
	router := httpAdapter.NewRouter(handler)

	// server
	server := httpserver.New(router, ":8080")

	go server.Run()
	waitForShutdown()
	server.Shutdown(context.Background())
}
