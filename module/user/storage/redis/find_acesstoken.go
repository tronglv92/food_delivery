package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
)

func (c *authUserCached) FindAccessToken(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*common.RedisToken, error) {
	var redisAccessToken *common.RedisToken

	accessToken := conditions[common.KeyRedisAccessToken].(string)
	userId := conditions["id"].(int)

	key := fmt.Sprintf(cacheKeyAT, userId, accessToken)

	err := c.cacheStore.Get(ctx, key, &redisAccessToken)

	if err != nil {
		return nil, err
	}

	return redisAccessToken, nil
}
