package managers

import "context"

type writerManager struct {
}

func NewWriterManager() Writer {
	return &writerManager{}
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
