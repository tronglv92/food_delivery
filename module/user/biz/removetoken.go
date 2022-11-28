package userbiz

import (
	"context"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/go-sdk/logger"
)

type RedisRemoveTokenStorage interface {
	GetKeys(ctx context.Context, conditions map[string]interface{}) ([]string, error)
	DelKeys(ctx context.Context, keys []string) error
}
type removeTokenBusiness struct {
	redisStorage RedisRemoveTokenStorage
}

func NewRemoveTokenBusiness(redisStorage RedisRemoveTokenStorage) *removeTokenBusiness {
	return &removeTokenBusiness{
		redisStorage: redisStorage,
	}
}
func (business *removeTokenBusiness) RemoveRedisToken(ctx context.Context, userId int) error {
	logger := logger.GetCurrent().GetLogger("module.user.biz.removetoken")
	keys, err := business.redisStorage.GetKeys(ctx, map[string]interface{}{"id": userId})

	if err != nil {
		return common.ErrInternal(err)
	}

	if keys == nil {
		return usermodel.ErrTokenNotFindInRedis
	}
	logger.Debugf("keys", keys)
	_ = business.redisStorage.DelKeys(ctx, keys)

	return nil
}
