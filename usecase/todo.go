package usecase

import (
	"github.com/akhr77/go-clean-architecture/domain/model"
	"github.com/akhr77/go-clean-architecture/domain/repository"
)

type TodoUsecase interface {
	Search(string) (todo []*model.Todo, err error)
	View() (todo []*model.Todo, err error)
	Add(*model.Todo) (err error)
	Edit(*model.Todo) (err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := TodoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

func (usecase *todoUsecase) Search(word string) ([]*model.Todo, error) {
	todo, err := usecase.todoRepo.Find(word)
	return todo, err
}

func (usecase *todoUsecase) View() ([]*model.Todo, error) {
	todo, err := usecase.todoRepo.FindAll()
	return todo, err
}

func (usecase *TodoUsecase) Add(todo *model.Todo) error {
	_, err := usecase.todoRepo.Create(todo)
	return err
}

func (usecase *todoUsecase) Edit(todo *model.Todo) error {
	_, err := usecase.todoRepo.Update(todo)
	return err
}
