package main

import (
	"feature-flag-poc/internal/config"
	"fmt"
	"log"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalln("failed to load config", err)
	}

	fmt.Println(config.Env)
}
