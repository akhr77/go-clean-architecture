package main

import (
	"fmt"

	"github.com/akhr77/go-clean-architecture/handler"
	"github.com/akhr77/go-clean-architecture/injector"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("server start")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
