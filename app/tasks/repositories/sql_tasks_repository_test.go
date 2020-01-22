package repositories

import (
	"context"
	"fance/app/tasks/domain"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
	"testing"
)

type sqlRepositorySuite struct {
	suite.Suite
	repository Repository
	ctx        context.Context
	id         string
}

func (suite *sqlRepositorySuite) SetupTest() {
	suite.ctx = context.Background()
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, _ := gorm.Open(mocket.DriverName, "connection_mock")
	suite.repository = NewSqlTasksRepository(db)
	suite.id = "id"
}

func TestRepository(t *testing.T) {
	suite.Run(t, &sqlRepositorySuite{})
}

var (
	task0 = domain.Task{
		Id:          "id",
		Title:       "as a developer",
		Description: "test cases",
		Status:      domain.ToDo,
	}
	task1 = domain.Task{
		Id: "id",
	}
	tasks0 = []domain.Task{task0}
)

func (suite *sqlRepositorySuite) TestSqlTasksRepository_GetAll() {
	tasksReply := getTasksDbReply()
	mocket.Catcher.Reset().NewMock().
		WithQuery(`SELECT * FROM "db_tasks"  WHERE (id is not null) ORDER BY id desc LIMIT 1 OFFSET 0`).
		WithReply(tasksReply)
	tasks, err := suite.repository.GetAll(suite.ctx, 1, 1)
	suite.NoError(err)
	suite.Equal(tasks0, tasks)
}

func getTasksDbReply() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":          task0.Id,
			"title":       task0.Title,
			"description": task0.Description,
			"status":      task0.Status,
		},
	}
}

func (suite *sqlRepositorySuite) TestSqlTasksRepository_Create_Success() {
	mocket.Catcher.Reset()
	result, err := suite.repository.Create(suite.ctx, &task0)
	suite.Nil(err)
	suite.Equal(&task0, result)
}

func (suite *sqlRepositorySuite) TestSqlTasksRepository_Create_Failed() {
	mocket.Catcher.Reset().NewMock().WithError(&mysql.MySQLError{Number: 1062, Message: "Error Entry"})
	result, err := suite.repository.Create(suite.ctx, &task0)
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *sqlRepositorySuite) TestSqlTasksRepository_Update() {
	mocket.Catcher.Reset()
	task1.Status = domain.Done
	err := suite.repository.Update(suite.ctx, &task1)
	suite.Nil(err)
}

func (suite *sqlRepositorySuite) TestSqlTasksRepository_Delete() {
	mocket.Catcher.Reset()
	err := suite.repository.Delete(suite.ctx, task0.Id)
	suite.Nil(err)
}
