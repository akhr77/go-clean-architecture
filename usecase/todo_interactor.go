package usecase

import "github.com/akhr77/go-clean-architecture/domain"

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Search(word string) (domain.Todos, error) {
	todos, err := interactor.TodoRepository.Find(word)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (interactor *TodoInteractor) View() (domain.Todos, error) {
	todos, err := interactor.TodoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (interactor *TodoInteractor) Add(todo domain.Todo) error {
	err := interactor.TodoRepository.Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *TodoInteractor) Edit(todo domain.Todo) error {
	err := interactor.TodoRepository.Update(todo)
	if err != nil {
		return err
	}
	return nil
}
