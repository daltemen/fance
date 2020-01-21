package managers

import "context"

type readerManager struct {
}

func NewReaderManager() Reader {
	return &readerManager{}
}

func (r *readerManager) RetrieveAll(ctx context.Context, page int, limit int) ([]TaskInfo, error) {
	panic("implement me")
}
