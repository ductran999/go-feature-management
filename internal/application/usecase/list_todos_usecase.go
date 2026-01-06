package usecase

import (
	"context"
	"feature-flag-poc/internal/application/port"
	"feature-flag-poc/internal/domain"
)

type listTodoUsecase struct {
	flags    port.FeatureFlag
	todoRepo port.TodoRepository
}

func NewListTodoUsecase(flags port.FeatureFlag, todoRepo port.TodoRepository) port.ListTodoUsecase {
	return &listTodoUsecase{
		flags:    flags,
		todoRepo: todoRepo,
	}
}

func (uc *listTodoUsecase) Execute(ctx context.Context) ([]domain.Todo, error) {
	if !uc.flags.IsEnabled("todos.enable_list_all") {
		return nil, ErrFeatureIsDisabled
	}

	todos, err := uc.todoRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
