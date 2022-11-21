package userstore

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mgoStore struct {
	client *mongo.Client
}

func NewMongoStore(client *mongo.Client) *mgoStore {
	return &mgoStore{client: client}
}
