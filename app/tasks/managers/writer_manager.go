package managers

import (
	"context"
	"fance/app/tasks/repositories"
)

type writerManager struct {
	Repository repositories.Repository
}

func NewWriterManager(repository repositories.Repository) Writer {
	return &writerManager{Repository: repository}
}

func (w *writerManager) Create(ctx context.Context, task *TaskInfo) (*TaskInfo, error) {
	panic("implement me")
}

func (w *writerManager) Update(ctx context.Context, task *TaskInfo) (*TaskInfo, error) {
	panic("implement me")
}

func (w *writerManager) Delete(ctx context.Context, taskID string) error {
	panic("implement me")
}
