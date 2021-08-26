package repositories

import (
	"goland-demo/api/utils/datastructs"
	"goland-demo/databases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository is the repository for users
type UserRepository struct {
	userCollection *mongo.Collection
}

// NewUserRepository is the constructor for the UserRepository struct
func NewUserRepository() (repo *UserRepository, err error) {
	dbSvr, err := databases.NewMongoSvr()
	if err != nil {
		return nil, err
	}
	userCollection := dbSvr.Collection("users")

	repo = &UserRepository{userCollection: userCollection}
	return repo, nil
}

func (repo *UserRepository) GetUsers(c *gin.Context) (users []*datastructs.User, err error) {
	cursor, err := repo.userCollection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) GetUser(c *gin.Context, id int32) (user *datastructs.User, err error) {
	err = repo.userCollection.FindOne(c, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
