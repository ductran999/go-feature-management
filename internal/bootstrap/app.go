package bootstrap

import (
	"context"
	httpAdapter "feature-flag-poc/internal/adapter/http"
	adapterPostgres "feature-flag-poc/internal/adapter/postgresql"
	adapterUnleash "feature-flag-poc/internal/adapter/unleash"
	"feature-flag-poc/internal/application/usecase"
	"feature-flag-poc/internal/config"
	"feature-flag-poc/internal/db/generated"
	"feature-flag-poc/internal/infra/db/postgresql"
	httpserver "feature-flag-poc/internal/server/http"
	"log"

	"github.com/Unleash/unleash-go-sdk/v5"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalln("failed to load config", err)
	}
	log.Println("[INFO] load config env successfully!")

	// init unleash
	if err := initUnleash(unleashConfig{
		AppName:    cfg.App.Name,
		BackendUrl: cfg.Unleash.BackendUrl,
		Token:      cfg.Unleash.Token,
		Env:        cfg.App.Environment,
	}); err != nil {
		log.Fatalln("failed to init unleash")
	}
	log.Println("[INFO] connect unleash successfully!")

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
	log.Println("[INFO] connect db successfully!")

	featureFlags := adapterUnleash.NewUnleashFeatureFlag()
	queries := generated.New(pool)
	repo := adapterPostgres.NewTodoRepository(queries)

	// usecase
	uc := usecase.NewListTodoUsecase(featureFlags, repo)

	// http adapter
	handler := httpAdapter.NewTodoHandler(uc)
	router := httpAdapter.NewRouter(handler)

	// server
	server := httpserver.New(router, ":8080")

	go server.Run()
	waitForShutdown()
	server.Shutdown(context.Background())
	if err := unleash.Close(); err != nil {
		log.Println("[WARN] got error when close unleash")
	}
}
