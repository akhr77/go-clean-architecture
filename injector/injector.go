package injector

import (
	"github.com/akhr77/go-clean-architecture/domain/repository"
	"github.com/akhr77/go-clean-architecture/handler"
	"github.com/akhr77/go-clean-architecture/infra"
	"github.com/akhr77/go-clean-architecture/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlhandler := InjectDB()
	return infra.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}
