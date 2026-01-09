# go-feature-management

[![Go Report Card](https://goreportcard.com/badge/github.com/ductran999/go-feature-management)](https://goreportcard.com/report/github.com/ductran999/go-feature-management)
[![Go](https://img.shields.io/badge/Go-1.25.1-blue?logo=go)](https://golang.org)
[![License](https://img.shields.io/github/license/ductran999/dbkit)](LICENSE)


A simple Proof of Concept (PoC) project to explore Unleash feature flag management using Go.
This project is not production-ready. Its goal is to understand how Unleash works, how feature flags are evaluated, and how it can be integrated into a Go service.

## 🎯 Purpose
- Learn how Unleash evaluates feature flags
- Experiment with Unleash Go client
- Understand feature toggles, strategies, and context
- Validate if Unleash fits real-world Go services

## 🧪 Scope (PoC)

- Basic feature flag enable/disable
- Context-based evaluation (user, environment)
- Simple Unleash client integration
- No persistence, no caching strategy optimization
- No advanced rollout strategies

## 📦 Tech Stack

- Go
- Unleash
- Unleash Go Client

## 🚀 Getting Started

### 1. Run Unleash (Local): see at [Official Document](https://github.com/Unleash/unleash)
```sh
git clone git@github.com:Unleash/unleash.git
cd unleash
docker compose up -d
```

Unleash UI: http://localhost:4242 (Default credentials: admin / unleash4all)

### 2. Install Dependencies
```bash
go mod tidy
```

### 4. Create env file
```
make init project
``` 
This will spawn 2 file `.env` and `./configs/config.yml`. Fill yours config. 

- `.env`: contains env for docker compose setup db
- `./configs/config.yml`: app configs

### 4. Set up local app
```bash
make setup
```

### 5. Migrate db schemas
I'm using [migrate](https://github.com/golang-migrate/migrate). You can simply connect to db and load the `.sql` file.
```bash
make migrate
```

### 6. Config flag
```go
// internal/application/usecase/list_todos_usecase.go
func (uc *listTodoUsecase) Execute(ctx context.Context) ([]domain.Todo, error) {
    // evaluate flag
	if !uc.flags.IsEnabled("todos.enable_list_all") {
		return nil, ErrFeatureIsDisabled
	}

	todos, err := uc.todoRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
```

Create a flag: `todos.enable_list_all` and disable it in development
![alt text](./assets/flags.png)

### 7. Result

```bash
curl --location 'localhost:8080/todos'
```

output:
```json
{
    "code": "INTERNAL_SERVER_ERROR",
    "message": "feature is disabled"
}
```

---

## 📄 License

This project is licensed under the [MIT](LICENSE) License.
