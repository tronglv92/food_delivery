package userstore

import "github.com/go-resty/resty/v2"

type userRestfulStore struct {
	client *resty.Client
}

func NewUserRestfulStore(client *resty.Client) *userRestfulStore {

	return &userRestfulStore{client: client}
}
