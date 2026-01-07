package main

import (
	bootstrap "feature-flag-poc/internal/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatalln(err)
	}
}
