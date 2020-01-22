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

type readerManagerSuite struct {
	suite.Suite
	repository *mocks.Repository
	ctx        context.Context
	reader     Reader
}

func (suite *readerManagerSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repository = new(mocks.Repository)
	suite.reader = NewReaderManager(suite.repository)
}

func TestReaderManager(t *testing.T) {
	suite.Run(t, &readerManagerSuite{})
}

var (
	tasks0 = []domain.Task{task0}
	tasksInfo0 = []TaskInfo{taskInfo0}
)

func (suite *readerManagerSuite) TestReaderManager_RetrieveAll_Success() {
	suite.repository.Mock.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(tasks0, nil).Once()
	result, err := suite.reader.RetrieveAll(suite.ctx, 1, 1)
	suite.NoError(err)
	suite.Equal(tasksInfo0, result)
}

func (suite *readerManagerSuite) TestReaderManager_RetrieveAll_Failed() {
	suite.repository.Mock.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()
	result, err := suite.reader.RetrieveAll(suite.ctx, 2, 1)
	suite.Error(err)
	suite.Nil(result)
}
