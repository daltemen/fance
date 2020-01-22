package managers

import "context"

type Reader interface {
	RetrieveAll(ctx context.Context, page int, limit int) ([]TaskInfo, error)
}
