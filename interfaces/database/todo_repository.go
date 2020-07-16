package database

import (
	"log"

	"github.com/akhr77/go-clean-architecture/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo TodoRepository) Find(word string) (domain.Todos, error) {
	rows, err := repo.Query("SELECT * FROM todos WHERE task LIKE ?", "%"+word+"%")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var todos domain.Todos
	for rows.Next() {
		todo := domain.Todo{}
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, todo)
	}
	return todos, nil
}

func (repo TodoRepository) FindAll() (domain.Todos, error) {
	rows, err := repo.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var todos domain.Todos
	for rows.Next() {
		todo := domain.Todo{}
		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, todo)
	}
	return todos, nil
}

func (repo *TodoRepository) Create(t domain.Todo) error {
	_, err := repo.Execute(
		"INSERT INTO todos (task, limit_date, status) VALUES (?, ?, ?)", t.Task, t.LimitDate, t.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TodoRepository) Update(t domain.Todo) error {
	_, err := repo.Execute("UPDATE todos SET task = ?, limit_date= ?, status = ?, where id = ?", t.Task, t.LimitDate, t.Status, t.ID)
	return err
}
