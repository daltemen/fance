package managers

import (
	"context"
	"errors"
	"fance/app/tasks/domain"
	"fance/app/tasks/repositories/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type writerManagerSuite struct {
	suite.Suite
	repository *mocks.Repository
	ctx        context.Context
	writer     Writer
}

func (suite *writerManagerSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repository = new(mocks.Repository)
	suite.writer = NewWriterManager(suite.repository)
}

func TestWriterManager(t *testing.T) {
	suite.Run(t, &writerManagerSuite{})
}

var (
	task0 = domain.Task{
		Id:          "id",
		Title:       "as a developer",
		Description: "test cases",
		Status:      domain.ToDo,
	}
	taskInfo0 = TaskInfo{
		Id:          task0.Id,
		Title:       task0.Title,
		Description: task0.Description,
		Status:      string(domain.ToDo),
	}
)

func (suite *writerManagerSuite) TestWriterManager_Create_Success() {
	suite.repository.Mock.On("Create", mock.Anything, mock.Anything).Return(&task0, nil).Once()
	result, err := suite.writer.Create(suite.ctx, &taskInfo0)
	suite.NoError(err)
	suite.Equal(&taskInfo0, result)
}

func (suite *writerManagerSuite) TestWriterManager_Create_Failed() {
	suite.repository.Mock.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()
	result, err := suite.writer.Create(suite.ctx, &taskInfo0)
	suite.Error(err)
	suite.Nil(result)
}

func (suite *writerManagerSuite) TestWriterManager_Update_Success() {
	suite.repository.Mock.On("Update", mock.Anything, mock.Anything).Return(nil).Once()
	result, err := suite.writer.Update(suite.ctx, &taskInfo0)
	suite.NoError(err)
	suite.Equal(&taskInfo0, result)
}

func (suite *writerManagerSuite) TestWriterManager_Update_Failed() {
	suite.repository.Mock.On("Update", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()
	result, err := suite.writer.Update(suite.ctx, &taskInfo0)
	suite.Error(err)
	suite.Nil(result)
}

func (suite *writerManagerSuite) TestWriterManager_Delete_Success() {
	suite.repository.Mock.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
	err := suite.writer.Delete(suite.ctx, taskInfo0.Id)
	suite.NoError(err)
}

func (suite *writerManagerSuite) TestWriterManager_Delete_Failed() {
	suite.repository.Mock.On("Delete", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()
	err := suite.writer.Delete(suite.ctx, taskInfo0.Id)
	suite.Error(err)
}
