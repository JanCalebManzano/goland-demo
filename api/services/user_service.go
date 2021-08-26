package services

import (
	"fmt"
	"goland-demo/api/repositories"
	"goland-demo/api/utils/datastructs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserService is the service for users
type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService() (srv *UserService, err error) {
	repo, err := repositories.NewUserRepository()
	if err != nil {
		return nil, err
	}

	srv = &UserService{
		repo: repo,
	}
	return srv, nil
}

func (srv *UserService) GetUsers(c *gin.Context) (users []*datastructs.User, err error) {
	return srv.repo.GetUsers(c)
}

func (srv *UserService) GetUser(c *gin.Context, sid string) (users *datastructs.User, err error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return nil, fmt.Errorf("Param with key 'id' should be an integer. ")
	}

	return srv.repo.GetUser(c, int32(id))
}
