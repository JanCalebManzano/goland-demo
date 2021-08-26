package services

import (
	"fmt"
	"goland-demo/api/repositories"
	"goland-demo/api/utils/datastructs"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (srv *TodoService) GetTodos(c *gin.Context) (todos []*datastructs.Todo, err error) {
	return srv.repo.GetTodos(c)
}

func (srv *TodoService) GetTodo(c *gin.Context, sid string) (todo *datastructs.Todo, err error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return nil, fmt.Errorf("Param with key 'id' should be an integer. ")
	}

	return srv.repo.GetTodo(c, int32(id))
}
