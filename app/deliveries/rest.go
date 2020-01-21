package deliveries

import (
	"fance/app/datasources"
	"fance/app/tasks/managers"
	"fance/app/tasks/repositories"
	"fance/app/tasks/rest"
	"os"

	"github.com/labstack/echo/v4"
)

func RunRestServer() {
	// -- init db
	db := datasources.ConnectDb()
	datasources.Migrate(db)
	// -- conn Interfaces
	repository := repositories.NewSqlTasksRepository(db)
	reader := managers.NewReaderManager(repository)
	writer := managers.NewWriterManager(repository)
	restMethods := rest.NewRest(reader, writer)
	// -- init rest server
	e := echo.New()
	rest.RegisterRoutes(e, restMethods)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
