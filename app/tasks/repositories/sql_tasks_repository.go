package repositories

import (
	"context"
	"fance/app/tasks/domain"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type sqlTasksRepository struct {
	Conn *gorm.DB
}

func NewSqlTasksRepository(conn *gorm.DB) Repository {
	return &sqlTasksRepository{Conn: conn}
}

func (s *sqlTasksRepository) GetAll(ctx context.Context, page int, limit int) ([]domain.Task, error) {
	var tasks []DbTask

	pagination.Paging(&pagination.Param{
		DB:      s.Conn.Where("id > ?", 0),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &tasks)

	return s.mapDbSliceToTask(tasks), nil
}

func (s *sqlTasksRepository) mapDbSliceToTask(tasks []DbTask) []domain.Task {
	result := make([]domain.Task, len(tasks))
	for i, t := range tasks {
		result[i] = domain.Task{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			Status:      domain.NewStatus(t.Status),
		}
	}
	return result
}

func (s *sqlTasksRepository) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	err := s.Conn.Create(&DbTask{Id: task.Id, Title: task.Title, Description: task.Description, Status: task.Status.String()}).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *sqlTasksRepository) Update(ctx context.Context, task *domain.Task) error {
	dbTask := &DbTask{Id: task.Id, Title: task.Title, Description: task.Description, Status: task.Status.String()}
	return s.Conn.Model(&dbTask).Update(dbTask).Error
}

func (s *sqlTasksRepository) Delete(ctx context.Context, taskID string) error {
	return s.Conn.Delete(&DbTask{Id: taskID}).Error
}
