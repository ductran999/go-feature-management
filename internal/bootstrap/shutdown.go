package bootstrap

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func waitForShutdown(
	appCtx context.Context,
	cancel context.CancelFunc,
	errChan chan error,
) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case sig := <-quit:
		log.Printf("[INFO] shutdown signal: %s", sig)
		cancel()

	case err := <-errChan:
		log.Printf("[ERROR] server error: %v", err)
		cancel()

	case <-appCtx.Done():
		log.Println("[INFO] context cancelled")
	}
}
