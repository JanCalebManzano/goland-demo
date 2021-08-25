package repositories

import (
	"goland-demo/api/utils/datastructs"
)

// TodoRepository is the repository for todos
type TodoRepository struct{}

func (repo *UserRepository) GetTodos() (todos []*datastructs.Todo, err error) {
	return []*datastructs.Todo{{
		ID:          0,
		UserID:      0,
		Content:     "Todo content",
		IsCompleted: false,
		CreatedAt:   nil,
		IsActive:    false,
	}}, nil
}

func (repo *UserRepository) GetTodo() (todo *datastructs.Todo, err error) {
	return &datastructs.Todo{
		ID:          0,
		UserID:      0,
		Content:     "Todo content",
		IsCompleted: false,
		CreatedAt:   nil,
		IsActive:    false,
	}, nil
}
