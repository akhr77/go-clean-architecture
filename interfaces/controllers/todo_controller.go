package controllers

import (
	"net/http"

	"github.com/akhr77/go-clean-architecture/domain"
	"github.com/akhr77/go-clean-architecture/interfaces/database"
	"github.com/akhr77/go-clean-architecture/usecase"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoHandler(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) View(c Context) error {
	models, err := controller.Interactor.View()
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return err
	}
	c.JSON(http.StatusOK, models)
	return nil
}

func (controller *TodoController) Search(c Context) error {
	word := c.QueryParam("word")
	models, err := controller.Interactor.Search(word)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return err
	}
	c.JSON(http.StatusOK, models)
	return nil
}

func (controller *TodoController) Add(c Context) error {
	todo := domain.Todo{}
	c.Bind(&todo)
	err := controller.Interactor.Add(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return err
	}
	c.JSON(http.StatusOK, nil)
	return nil
}

func (controller *TodoController) Edit(c Context) error {
	todo := domain.Todo{}
	c.Bind(todo)
	err := controller.Interactor.Edit(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return err
	}
	c.JSON(http.StatusOK, nil)
	return nil
}
