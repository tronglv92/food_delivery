package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
)

func (s *mgoStore) CreateUser(ctxs context.Context, data *usermodel.UserCreate) error {

	collection := s.client.Database(common.DBMongoName).Collection(common.UsersCollection)
	result, err := collection.InsertOne(ctxs, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result ", result)
	return nil
}
