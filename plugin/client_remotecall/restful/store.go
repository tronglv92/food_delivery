package remoterestful

import (
	"github.com/go-resty/resty/v2"
)

type userRestfulStore struct {
	client     *resty.Client
	serviceURL string
}

func NewUserRestfulStore(client *resty.Client, serviceURL string) *userRestfulStore {

	return &userRestfulStore{client: client, serviceURL: serviceURL}
}
