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

func (ctrl *UserController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.srv.GetUsers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (ctrl *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		res, err := ctrl.srv.GetUser(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
