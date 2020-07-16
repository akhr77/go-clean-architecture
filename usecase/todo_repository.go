package usecase

import (
	"github.com/akhr77/go-clean-architecture/domain"
)

type TodoRepository interface {
	Find(string) (todo domain.Todos, err error)
	FindAll() (todo domain.Todos, err error)
	Create(domain.Todo) (err error)
	Update(domain.Todo) (err error)
}
