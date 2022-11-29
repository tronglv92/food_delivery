package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/plugin/tokenprovider"
	"time"
)

func (c *authUserCached) WLSaveTokens(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) error {
	userId := conditions["id"].(int)
	accessToken := conditions[common.KeyRedisAccessToken].(tokenprovider.Token)
	refreshToken := conditions[common.KeyRedisRefreshToken].(tokenprovider.Token)
	keyAT := fmt.Sprintf(common.CacheWLKeyAT, userId, accessToken.GetToken())

	keyRT := fmt.Sprintf(common.CacheWLKeyRT, userId, refreshToken.GetToken())

	redisAccessToken := common.RedisToken{
		Token:   accessToken.GetToken(),
		Created: time.Now(),
		Expiry:  accessToken.GetExpire(),
	}

	redisRefreshToken := common.RedisToken{
		Token:   refreshToken.GetToken(),
		Created: time.Now(),
		Expiry:  refreshToken.GetExpire(),
	}
	_ = c.cacheStore.Set(ctx, keyAT, redisAccessToken, time.Duration(accessToken.GetExpire()*int(time.Second)))
	_ = c.cacheStore.Set(ctx, keyRT, redisRefreshToken, time.Duration(refreshToken.GetExpire()*int(time.Second)))

	return nil
}
