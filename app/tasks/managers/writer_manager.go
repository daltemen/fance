package managers

import (
	"context"
	"fance/app/tasks/domain"
	"fance/app/tasks/repositories"
)

type writerManager struct {
	repo repositories.Repository
}

func NewWriterManager(repository repositories.Repository) Writer {
	return &writerManager{repo: repository}
}

func (w *writerManager) Create(ctx context.Context, task *TaskInfo) (*TaskInfo, error) {
	n := domain.NewTask(task.Title, task.Description, domain.NewStatus(task.Status))
	created, err := w.repo.Create(ctx, n)
	if err != nil {
		return nil, err
	}
	return w.mapTaskToTaskInfo(created), nil
}

func (w *writerManager) mapTaskToTaskInfo(task *domain.Task) *TaskInfo {
	return &TaskInfo{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status.String(),
	}
}

func (w *writerManager) Update(ctx context.Context, task *TaskInfo) (*TaskInfo, error) {
	if err := w.repo.Update(ctx, w.mapTaskInfoToDomain(task)); err != nil {
		return nil, err
	}
	return task, nil
}

func (w *writerManager) mapTaskInfoToDomain(task *TaskInfo) *domain.Task {
	return &domain.Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      domain.NewStatus(task.Status),
	}
}

func (w *writerManager) Delete(ctx context.Context, taskID string) error {
	if err := w.repo.Delete(ctx, taskID); err != nil {
		return err
	}
	return nil
}
