package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/plugin/tokenprovider"
	"time"
)

const cacheKeyAT = "user:%d:%v"
const cacheKeyRT = "user:refreshtoken:%d"

func (c *authUserCached) SaveToken(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) error {
	userId := conditions["id"].(int)
	accessToken := conditions["access_token"].(tokenprovider.Token)
	// refreshToken := conditions["refresh_token"].(tokenprovider.Token)
	keyAT := fmt.Sprintf(cacheKeyAT, userId, accessToken.GetToken())
	fmt.Printf("keyAT ", keyAT)
	// keyRT := fmt.Sprintf(cacheKeyRT, userId)

	redisAccessToken := common.RedisToken{
		Token:   accessToken.GetToken(),
		Created: time.Now(),
		Expiry:  accessToken.GetExpire(),
	}
	_ = c.cacheStore.Set(ctx, keyAT, redisAccessToken, time.Duration(accessToken.GetExpire()*int(time.Second)))
	// _ = c.cacheStore.Set(ctx, keyRT, refreshToken.GetToken(), time.Duration(refreshToken.GetExpire()*int(time.Second)))

	return nil
}
