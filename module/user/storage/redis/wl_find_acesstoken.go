package userstore

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/common"

	"github.com/go-redis/cache/v8"
)

func (c *authUserCached) WLFindAccessToken(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*common.RedisToken, error) {
	var redisAccessToken *common.RedisToken

	accessToken := conditions[common.KeyRedisAccessToken].(string)
	userId := conditions["id"].(int)
	key := fmt.Sprintf(common.CacheWLKeyAT, userId, accessToken)

	err := c.cacheStore.Get(ctx, key, &redisAccessToken)

	if err != nil {
		if err == cache.ErrCacheMiss {
			return nil, errors.New("AccessToken don't exist in whitelist")
		}
		return nil, err
	}

	return redisAccessToken, nil
}
