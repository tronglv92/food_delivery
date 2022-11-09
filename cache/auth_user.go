package cache

import (
	"context"
	"fmt"
	usermodel "food_delivery/module/user/model"
	"sync"
	"time"
)

const cacheKey = "user:%d"

type UserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type authUserCached struct {
	realStore  UserStore
	cacheStore Cache
	once       *sync.Once
}

func NewAuthUserCache(realStore UserStore, cacheStore Cache) *authUserCached {
	return &authUserCached{
		realStore:  realStore,
		cacheStore: cacheStore,
		once:       new(sync.Once),
	}
}
func (c *authUserCached) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	var user usermodel.User

	userId := conditions["id"].(int)
	key := fmt.Sprintf(cacheKey, userId)

	_ = c.cacheStore.Get(ctx, key, &user)

	if user.Id > 0 {
		return &user, nil
	}

	if user.Id == 0 {
		var err error

		c.once.Do(func() {
			//log.Println("get user from DB")
			u, errDB := c.realStore.FindUser(ctx, conditions)

			if err != nil {
				err = errDB
			} else {
				user = *u
				_ = c.cacheStore.Set(ctx, key, u, time.Hour*2)
			}
		})
	}

	_ = c.cacheStore.Get(ctx, key, &user)

	return &user, nil
}
