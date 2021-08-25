package controllers

import (
	"goland-demo/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	srv *services.TodoService
}

func NewTodoController() (ctrl *TodoController, err error) {
	srv, err := services.NewTodoService()
	if err != nil {
		return nil, err
	}

	ctrl = &TodoController{srv: srv}
	return ctrl, nil
}

func (ctrl *TodoController) GetTodos(c *gin.Context) {
	res, err := ctrl.srv.GetTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl *TodoController) GetTodo(c *gin.Context) {
	res, err := ctrl.srv.GetTodo()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
