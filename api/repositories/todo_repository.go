package repositories

import (
	"goland-demo/api/utils/datastructs"
	"goland-demo/databases"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

// TodoRepository is the repository for todos
type TodoRepository struct {
	todoCollection *mongo.Collection
}

// NewTodoRepository is the constructor for the TodoRepository struct
func NewTodoRepository() (repo *TodoRepository, err error) {
	dbSvr, err := databases.NewMongoSvr()
	if err != nil {
		return nil, err
	}
	todoCollection := dbSvr.Collection("todos")

	repo = &TodoRepository{todoCollection: todoCollection}
	return repo, nil
}

func (repo *TodoRepository) GetTodos(c *gin.Context) (todos []*datastructs.Todo, err error) {
	cursor, err := repo.todoCollection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (repo *TodoRepository) GetTodo(c *gin.Context, id int32) (todo *datastructs.Todo, err error) {
	err = repo.todoCollection.FindOne(c, bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
