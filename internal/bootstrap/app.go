package bootstrap

import (
	"context"
	"errors"
	httpAdapter "feature-flag-poc/internal/adapter/http"
	adapterPostgres "feature-flag-poc/internal/adapter/postgresql"
	adapterUnleash "feature-flag-poc/internal/adapter/unleash"
	"feature-flag-poc/internal/application/usecase"
	"feature-flag-poc/internal/config"
	"feature-flag-poc/internal/db/generated"
	"feature-flag-poc/internal/infra/db/postgresql"
	httpserver "feature-flag-poc/internal/server/http"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Unleash/unleash-go-sdk/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	pool   *pgxpool.Pool
	server *httpserver.Server
	once   sync.Once
}

func (app *App) Close() {
	app.once.Do(func() {
		log.Println("[INFO] shutting down application...")

		if app.server != nil {
			log.Println("[INFO] shutting down api server...")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := app.server.Shutdown(ctx); err != nil {
				log.Printf("[WARN] api server shutdown error: %v", err)
			} else {
				log.Println("[INFO] api server stopped successfully")
			}
		}

		if app.pool != nil {
			log.Println("[INFO] closing database connection pool...")
			app.pool.Close()
			log.Println("[INFO] database pool closed successfully")
		}

		if err := unleash.Close(); err != nil {
			log.Printf("[WARN] unleash close error: %v", err)
		} else {
			log.Println("[INFO] unleash closed successfully")
		}

		log.Println("[INFO] application shutdown completed")
	})
}

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := App{}
	defer app.Close()

	cfg, err := config.LoadEnv()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	log.Println("[INFO] load config env successfully!")

	// init unleash
	if err := initUnleash(unleashConfig{
		AppName:    cfg.App.Name,
		BackendUrl: cfg.Unleash.BackendUrl,
		Token:      cfg.Unleash.Token,
		Env:        cfg.App.Environment,
	}); err != nil {
		return fmt.Errorf("failed to init unleash: %w", err)
	}
	log.Println("[INFO] connect unleash successfully!")

	//  Init DB
	app.pool, err = postgresql.New(ctx, postgresql.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Pass,
		DBName:   cfg.DB.DatName,
		SSLMode:  "disable",
	})
	if err != nil {
		return fmt.Errorf("failed to connect db: %w", err)
	}
	log.Println("[INFO] connect db successfully!")

	featureFlags := adapterUnleash.NewUnleashFeatureFlag()
	queries := generated.New(app.pool)
	repo := adapterPostgres.NewTodoRepository(queries)

	// usecase
	uc := usecase.NewListTodoUsecase(featureFlags, repo)

	// http adapter
	handler := httpAdapter.NewTodoHandler(uc)
	router := httpAdapter.NewRouter(handler)

	// server
	app.server = httpserver.New(router, ":8080")

	errChan := make(chan error, 1)
	go func() {
		if err := app.server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	waitForShutdown(ctx, cancel, errChan)

	return nil
}
