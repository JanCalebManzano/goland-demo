package databases

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoSvr struct {
	db *mongo.Database
}

var (
	mongoDBInit     sync.Once
	mongoDBInstance *MongoSvr
)

func NewMongoSvr() (*MongoSvr, error) {
	var (
		client *mongo.Client
		db     *mongo.Database
		err    error
	)

	mongoDBInit.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s",
			os.Getenv("DB_MONGO_HOST"), os.Getenv("DB_MONGO_PORT"))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			return
		}

		// Ping the primary
		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			return
		}

		db = client.Database(os.Getenv("DB_MONGO_NAME"))

		mongoDBInstance = &MongoSvr{
			db: db,
		}
	})

	if err != nil {
		return nil, err
	}

	return mongoDBInstance, nil
}

func (dbSvr *MongoSvr) Collection(name string) *mongo.Collection {
	return dbSvr.db.Collection(name)
}
