package userstore

import (
	"context"
	"fmt"
	usermodel "food_delivery/module/user/model"

	"time"
)



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
