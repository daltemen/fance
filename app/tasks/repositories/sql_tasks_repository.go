package repositories

import (
	"context"
	"fance/app/tasks"
)

type sqlTasksRepository struct{

}

func (s *sqlTasksRepository) GetAll(ctx context.Context, page int, limit int) ([]tasks.Task, error) {
	panic("implement me")
}

func (s *sqlTasksRepository) Create(ctx context.Context, task *tasks.Task) (*tasks.Task, error) {
	panic("implement me")
}

func (s *sqlTasksRepository) Update(ctx context.Context, task *tasks.Task) error {
	panic("implement me")
}

func (s *sqlTasksRepository) Delete(ctx context.Context, taskID string) error {
	panic("implement me")
}
