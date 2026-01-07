package postgresql

import (
	"context"
	"feature-flag-poc/internal/application/port"
	"feature-flag-poc/internal/db/generated"
	"feature-flag-poc/internal/domain"
)

type TodoRepository struct {
	q *generated.Queries
}

func NewTodoRepository(q *generated.Queries) port.TodoRepository {
	return &TodoRepository{q: q}
}

func (r *TodoRepository) List(ctx context.Context) ([]domain.Todo, error) {
	rows, err := r.q.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Todo, 0, len(rows))
	for _, row := range rows {
		result = append(result, domain.Todo{
			ID:        row.ID,
			Title:     row.Title,
			Status:    string(row.Status),
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		})
	}

	return result, nil
}
