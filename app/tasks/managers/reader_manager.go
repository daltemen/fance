package managers

import (
	"context"
	"fance/app/tasks/domain"
	"fance/app/tasks/repositories"
)

type readerManager struct {
	repo repositories.Repository
}

func NewReaderManager(repository repositories.Repository) Reader {
	return &readerManager{repo: repository}
}

func (r *readerManager) RetrieveAll(ctx context.Context, page int, limit int) ([]TaskInfo, error) {
	result, err := r.repo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return r.mapDomainToTaskInfo(result), nil
}

func (r *readerManager) mapDomainToTaskInfo(tasks []domain.Task) []TaskInfo {
	result := make([]TaskInfo, len(tasks))
	for i, task := range tasks {
		result[i] = TaskInfo{
			Id:          task.Id,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status.String(),
		}
	}
	return result
}
