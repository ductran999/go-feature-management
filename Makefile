-include .env
export $(shell sed 's/=.*//' .env 2>/dev/null)

default: help

help: ## Show help for each of the Makefile commands
	@awk 'BEGIN \
		{FS = ":.*##"; printf "Usage: make ${cyan}<command>\n${white}Commands:\n"} \
		/^[a-zA-Z_-]+:.*?##/ \
		{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' \
		$(MAKEFILE_LIST)

.PHONY: tidy
tidy: ## Tidy up the go.mod
	go mod tidy

.PHONY: lint
lint: ## Run linters
	golangci-lint run --timeout 10m --config .golangci.yml

.PHONY: 
build: ## Build image
	@docker build -t ducnta99/go-feature-flags:${APP_VERSION} .

.PHONY: setup
setup: ## Make setup
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "Copied .env.example -> .env"; \
	else \
		echo ".env already exists, skipping copy."; \
	fi

	@if [ ! -f ./configs/dev/config.yml ]; then \
		cp ./configs/dev/example.yml ./configs/dev/config.yml; \
		echo "Copied ./configs/dev/example.yml -> ./configs/dev/config.yml"; \
	else \
		echo "./configs/dev/config.yml already exists, skipping copy."; \
	fi

	@if [ ! -f ./configs/prod/config.yml ]; then \
		cp ./configs/prod/example.yml ./configs/prod/config.yml; \
		echo "Copied ./configs/prod/example.yml -> ./configs/prod/config.yml"; \
	else \
		echo "./configs/prod/config.yml already exists, skipping copy."; \
	fi

	@docker compose up -d


.PHONY: cleanup
cleanup: ## Cleanup demo
	@docker compose down

.PHONY: sqlc-gen
sqlc-gen: ## generate code
	@sqlc generate -f ./internal/db/sqlc.yaml

.PHONY: migrate
migrate: ## migrate database schema
	migrate \
	-path migrations \
  	-database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_DATABASE}?sslmode=disable" \
  	up

.PHONY: run
run: ## run demo
	go run ./cmd/main.go
