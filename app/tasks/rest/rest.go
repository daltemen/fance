package rest

import (
	"fance/app/tasks/managers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Rest struct {
	Reader managers.Reader
	Writer managers.Writer
}

func NewRest(reader managers.Reader, writer managers.Writer) *Rest {
	return &Rest{Reader: reader, Writer: writer}
}

func (r *Rest) GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "Get, all!")
}

func (r *Rest) PostTask(c echo.Context) error {
	return c.String(http.StatusOK, "post, task!")
}

func (r *Rest) PutTask(c echo.Context) error {
	return c.String(http.StatusOK, "put task!")
}

func (r *Rest) DeleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "delete task!")
}
