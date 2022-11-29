package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
)

func (c *authUserCached) FindRefreshToken(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*common.RedisToken, error) {
	var redisRefreshToken *common.RedisToken

	refreshToken := conditions[common.KeyRedisRefreshToken].(string)
	userId := conditions["id"].(int)

	key := fmt.Sprintf(cacheKeyRT, userId, refreshToken)

	err := c.cacheStore.Get(ctx, key, &redisRefreshToken)

	if err != nil {
		return nil, err
	}

	return redisRefreshToken, nil
}
