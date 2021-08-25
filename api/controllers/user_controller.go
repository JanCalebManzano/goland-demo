package controllers

import (
	"goland-demo/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv *services.UserService
}

func NewUserController() (ctrl *UserController, err error) {
	srv, err := services.NewUserService()
	if err != nil {
		return nil, err
	}

	ctrl = &UserController{srv: srv}
	return ctrl, nil
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	res, err := ctrl.srv.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	res, err := ctrl.srv.GetUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
