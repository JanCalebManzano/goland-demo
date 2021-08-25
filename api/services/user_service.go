package services

import (
	"goland-demo/api/repositories"
	"goland-demo/api/utils/datastructs"
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

func (srv *UserService) GetUsers() (users []*datastructs.User, err error) {
	return srv.repo.GetUsers()
}

func (srv *UserService) GetUser() (users *datastructs.User, err error) {
	return srv.repo.GetUser()
}
