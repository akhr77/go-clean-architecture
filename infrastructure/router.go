package infrastructure

import (
	"github.com/akhr77/go-clean-architecture/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Router *echo.Echo

func Run() {
	// Echo instance
	e := echo.New()

	todoController := controllers.NewTodoHandler(NewSqlHandler())

	// MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error { return todoController.View(c) })
	e.GET("/search", func(c echo.Context) error { return todoController.Search(c) })
	e.POST("/todoCreate", func(c echo.Context) error { return todoController.Add(c) })
	e.POST("/todoEdit", func(c echo.Context) error { return todoController.Edit(c) })

	e.Logger.Fatal(e.Start(":8080"))

}
