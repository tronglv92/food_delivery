package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
)

func (c *authUserCached) BLFindTokens(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) bool {
	userId := conditions["id"].(int)

	refreshToken := conditions[common.KeyRedisRefreshToken].(string)

	keyRT := fmt.Sprintf(common.CacheBLKeyRT, userId, refreshToken)

	var redisRefreshToken *BlackToken

	err := c.cacheStore.Get(ctx, keyRT, &redisRefreshToken)
	fmt.Printf("keyRT %v", keyRT)
	fmt.Printf("redisRefreshToken %v", redisRefreshToken)
	fmt.Printf("err %v", err)

	return redisRefreshToken != nil
}
