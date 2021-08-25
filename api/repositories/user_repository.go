package repositories

import (
	"goland-demo/api/utils/datastructs"
)

// UserRepository is the repository for users
type UserRepository struct{}

func (repo *UserRepository) GetUsers() (users []*datastructs.User, err error) {
	return []*datastructs.User{{
		ID:           0,
		Username:     "jan01",
		FullName:     "Jan Caleb Salvador Manzano",
		MobileNumber: "+639123456789",
		EmailAddress: "jan01@simplexi.com.ph",
		CreatedAt:    nil,
		IsActive:     false,
	}}, nil
}

func (repo *UserRepository) GetUser() (user *datastructs.User, err error) {
	return &datastructs.User{
		ID:           0,
		Username:     "jan01",
		FullName:     "Jan Caleb Salvador Manzano",
		MobileNumber: "+639123456789",
		EmailAddress: "jan01@simplexi.com.ph",
		CreatedAt:    nil,
		IsActive:     false,
	}, nil
}
