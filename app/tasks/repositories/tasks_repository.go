package repositories

import (
	"context"
	"fance/app/tasks/domain"
)

type Repository interface {
	GetAll(ctx context.Context, page int, limit int) ([]domain.Task, error)
	Create(ctx context.Context, task *domain.Task) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, taskID string) error
}
