package userstore

import (
	"context"
	"fmt"
	"food_delivery/common"
	"math/rand"
	"time"
)

type BlackToken struct {
	Token string `json:"token"`

	//Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (c *authUserCached) BLSaveTokens(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) error {
	userId := conditions["id"].(int)

	refreshToken := conditions[common.KeyRedisRefreshToken].(string)

	keyRT := fmt.Sprintf(common.CacheBLKeyRT, userId, refreshToken)

	ttlToken := common.AddedBlackListDuration + time.Duration(rand.Int31n(100))*time.Second

	fmt.Printf("ttlToken %v", ttlToken)

	redisRefreshToken := BlackToken{
		Token: refreshToken,
	}
	_ = c.cacheStore.Set(ctx, keyRT, redisRefreshToken, ttlToken)

	return nil
}
