package managers

import (
	"context"
	"fance/app/tasks/repositories"
)

type readerManager struct {
	Repository repositories.Repository
}

func NewReaderManager(repository repositories.Repository) Reader {
	return &readerManager{Repository: repository}
}

func (r *readerManager) RetrieveAll(ctx context.Context, page int, limit int) ([]TaskInfo, error) {
	panic("implement me")
}
