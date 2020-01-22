package managers

import "context"

type Writer interface {
	Create(ctx context.Context, task *TaskInfo) (*TaskInfo, error)
	Update(ctx context.Context, task *TaskInfo) (*TaskInfo, error)
	Delete(ctx context.Context, taskID string) error
}
