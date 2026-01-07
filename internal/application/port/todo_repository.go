package port

import (
	"context"
	"feature-flag-poc/internal/domain"
)

type TodoRepository interface {
	List(ctx context.Context) ([]domain.Todo, error)
}
