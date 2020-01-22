package rest

import (
	"fance/app/tasks/managers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Rest struct {
	reader managers.Reader
	writer managers.Writer
}

func NewRest(reader managers.Reader, writer managers.Writer) *Rest {
	return &Rest{reader: reader, writer: writer}
}

func (r *Rest) GetAll(c echo.Context) error {
	page, limit := r.getPageAndLimit(c)
	tasks, err := r.reader.RetrieveAll(c.Request().Context(), page, limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r.mapTasksInfoToGetAll(tasks, page, limit))
}

func (r *Rest) mapTasksInfoToGetAll(taskInfo []managers.TaskInfo, page, limit int) TasksGetAll {
	result := TasksGetAll{
		Page:  page,
		Limit: limit,
	}
	tasks := make([]TaskResponse, len(taskInfo))
	for i, t := range taskInfo {
		tasks[i] = *r.mapTaskInfoToResponse(&t)
	}
	result.Tasks = tasks
	return result
}

func (r *Rest) getPageAndLimit(c echo.Context) (int, int) {
	p := c.QueryParam("page")
	l := c.QueryParam("limit")
	if p == "" || l == "" {
		return 1, 10
	}
	page, _ := strconv.Atoi(p)
	limit, _ := strconv.Atoi(l)
	return page, limit
}

func (r *Rest) PostTask(c echo.Context) error {
	t := new(TaskRequest)
	if err := c.Bind(t); err != nil {
		return err
	}
	task, err := r.writer.Create(c.Request().Context(), r.mapRequestToTaskInfo(t))
	if err != nil {
		return c.JSON(http.StatusConflict, ErrorRest{Msg: err.Error()})
	}
	return c.JSON(http.StatusCreated, r.mapTaskInfoToResponse(task))
}

func (r *Rest) PutTask(c echo.Context) error {
	t := new(TaskRequest)
	if err := c.Bind(t); err != nil {
		return err
	}
	info := r.mapRequestToTaskInfo(t)
	info.Id = c.Param("id")
	task, err := r.writer.Update(c.Request().Context(), info)
	if err != nil {
		return c.JSON(http.StatusConflict, ErrorRest{Msg: err.Error()})
	}
	return c.JSON(http.StatusOK, r.mapTaskInfoToResponse(task))
}

func (r *Rest) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := r.writer.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusConflict, ErrorRest{Msg: err.Error()})
	}
	return c.JSON(http.StatusNoContent, "")
}

func (r *Rest) mapRequestToTaskInfo(request *TaskRequest) *managers.TaskInfo {
	return &managers.TaskInfo{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}
}

func (r *Rest) mapTaskInfoToResponse(task *managers.TaskInfo) *TaskResponse {
	return &TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
}
