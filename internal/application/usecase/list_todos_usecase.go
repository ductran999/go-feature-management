package usecase

import (
	"context"
	"feature-flag-poc/internal/application/port"
	"feature-flag-poc/internal/domain"
)

type listTodoUsecase struct {
	todoRepo port.TodoRepository
}

func NewListTodoUsecase(todoRepo port.TodoRepository) port.ListTodoUsecase {
	return &listTodoUsecase{
		todoRepo: todoRepo,
	}
}

func (uc *listTodoUsecase) Execute(ctx context.Context) ([]domain.Todo, error) {
	todos, err := uc.todoRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
