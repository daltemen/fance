package repositories

import (
	"context"
	"fance/app/tasks"
)

type Repository interface {
	GetAll(ctx context.Context, page int, limit int) ([]tasks.Task, error)
	Create(ctx context.Context, task *tasks.Task) (*tasks.Task, error)
	Update(ctx context.Context, task *tasks.Task) error
	Delete(ctx context.Context, taskID string) error
}
