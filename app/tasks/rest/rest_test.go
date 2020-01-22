package rest

import (
	"context"
	"fance/app/tasks/managers"
	"fance/app/tasks/managers/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type restSuite struct {
	suite.Suite
	ctx    context.Context
	reader *mocks.Reader
	writer *mocks.Writer
	rest   *Rest
}

func (suite *restSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.reader = new(mocks.Reader)
	suite.writer = new(mocks.Writer)
	suite.rest = NewRest(suite.reader, suite.writer)
}

func TestRestSuite(t *testing.T) {
	suite.Run(t, &restSuite{})
}

var (
	taskInfo0 = managers.TaskInfo{
		Id:          "d1b9e7b2-923d-43b2-a1b0-63a96a02663f",
		Title:       "My Title",
		Description: "As a PM...",
		Status:      "DOING",
	}
	tasksInfo0 = []managers.TaskInfo{taskInfo0}

	getAllJSON = `{"tasks":[{"id":"d1b9e7b2-923d-43b2-a1b0-63a96a02663f","title":"My Title","description":"As a PM...","status":"DOING"}],"page":1,"limit":10}
`
	taskInfo1 = managers.TaskInfo{
		Id:          "d1b9e7b2-923d-43b2-a1b0-63a96a02663f",
		Title:       "Unit Tests",
		Description: "As a Android developer...",
		Status:      "TODO",
	}
	taskJSON   = `{ "title": "Unit Tests", "description": "As a Android developer...", "status": "TODO" }`

)

func (suite *restSuite) TestRest_GetAll() {
	suite.reader.Mock.On("RetrieveAll", mock.Anything, mock.Anything, mock.Anything).Return(tasksInfo0, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.GetAll(c)
	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
	suite.Equal(getAllJSON, rec.Body.String())
}


func (suite *restSuite) TestRest_PostTask() {
	suite.writer.Mock.On("Create", mock.Anything, mock.Anything).Return(&taskInfo1, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", strings.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.PostTask(c)
	expected := `{"id":"d1b9e7b2-923d-43b2-a1b0-63a96a02663f","title":"Unit Tests","description":"As a Android developer...","status":"TODO"}
`
	suite.NoError(err)
	suite.Equal(http.StatusCreated, rec.Code)
	suite.Equal(expected, rec.Body.String())
}

func (suite *restSuite) TestRest_PutTask() {
	suite.writer.Mock.On("Update", mock.Anything, mock.Anything).Return(&taskInfo1, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/tasks/" + taskInfo1.Id, strings.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.PutTask(c)
	expected := `{"id":"d1b9e7b2-923d-43b2-a1b0-63a96a02663f","title":"Unit Tests","description":"As a Android developer...","status":"TODO"}
`
	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
	suite.Equal(expected, rec.Body.String())
}

func (suite *restSuite) TestRest_DeleteTask() {
	suite.writer.Mock.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/tasks/" + taskInfo1.Id, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.DeleteTask(c)
	suite.NoError(err)
	suite.Equal(http.StatusNoContent, rec.Code)
}
