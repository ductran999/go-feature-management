package bootstrap

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func waitForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGTERM, // docker stop, k8s terminate
	)

	sig := <-quit
	log.Printf("shutdown signal received: %s", sig.String())
}
