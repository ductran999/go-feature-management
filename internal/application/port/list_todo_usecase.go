package port

import (
	"context"
	"feature-flag-poc/internal/domain"
)

type ListTodoUsecase interface {
	Execute(ctx context.Context) ([]domain.Todo, error)
}
