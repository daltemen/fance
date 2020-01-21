package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type EchoRestType func(c echo.Context) error

func RegisterRoutes(e *echo.Echo, rest *Rest) {
	e.GET("/health", Health)
	e.GET("/api/v1/tasks", rest.GetAll)
	e.POST("/api/v1/tasks", rest.PostTask)
	e.PUT("/api/v1/tasks/:id", rest.PutTask)
	e.DELETE("/api/v1/tasks/:id", rest.DeleteTask)
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
