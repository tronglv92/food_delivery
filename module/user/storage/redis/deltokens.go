package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
)

func (c *authUserCached) DelTokens(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) error {
	userId := conditions["id"].(int)
	accessToken := conditions[common.KeyRedisAccessToken].(string)
	refreshToken := conditions[common.KeyRedisRefreshToken].(string)

	keyAT := fmt.Sprintf(cacheKeyAT, userId, accessToken)
	keyRT := fmt.Sprintf(cacheKeyRT, userId, refreshToken)

	_ = c.cacheStore.Delete(ctx, keyAT)
	_ = c.cacheStore.Delete(ctx, keyRT)

	return nil
}
