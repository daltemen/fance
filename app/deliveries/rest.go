package deliveries

import (
	"fance/app/datasources"
	"fance/app/tasks/managers"
	"fance/app/tasks/rest"
	"os"

	"github.com/labstack/echo/v4"
)

func RunRestServer() {
	// -- init db
	db := datasources.ConnectDb()
	datasources.Migrate(db)
	// -- conn Interfaces
	reader := managers.NewReaderManager()
	writer := managers.NewWriterManager()
	restMethods := rest.NewRest(reader, writer)
	// -- init rest server
	e := echo.New()
	rest.RegisterRoutes(e, restMethods)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
