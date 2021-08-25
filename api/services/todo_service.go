package services

import (
	"goland-demo/api/repositories"
	"goland-demo/api/utils/datastructs"
)

// TodoService is the service for todos
type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService() (srv *TodoService, err error) {
	repo, err := repositories.NewTodoRepository()
	if err != nil {
		return nil, err
	}

	srv = &TodoService{
		repo: repo,
	}
	return srv, nil
}

func (srv *TodoService) GetTodos() (todos []*datastructs.Todo, err error) {
	return srv.repo.GetTodos()
}

func (srv *TodoService) GetTodo() (todo *datastructs.Todo, err error) {
	return srv.repo.GetTodo()
}
