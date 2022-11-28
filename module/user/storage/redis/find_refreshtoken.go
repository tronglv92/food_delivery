package userstore

import (
	"context"
	"fmt"
	"food_delivery/plugin/tokenprovider"
)

func (c *authUserCached) FindRefreshToken(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*tokenprovider.Token, error) {
	var refreshToken tokenprovider.Token

	userId := conditions["id"].(int)
	key := fmt.Sprintf(cacheKeyRT, userId)

	_ = c.cacheStore.Get(ctx, key, &refreshToken)

	return &refreshToken, nil
}
