package userstore

import (
	"context"

	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/storage/sdkredis"
	"sync"
)

type UserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type authUserCached struct {
	realStore  UserStore
	cacheStore sdkredis.Cache
	once       *sync.Once
}

func NewAuthUserCache(realStore UserStore, cacheStore sdkredis.Cache) *authUserCached {
	return &authUserCached{
		realStore:  realStore,
		cacheStore: cacheStore,
		once:       new(sync.Once),
	}
}
