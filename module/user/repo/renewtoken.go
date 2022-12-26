package userrepo

import (
	"context"
	"time"

	"food_delivery/common"

	usermodel "food_delivery/module/user/model"

	"food_delivery/plugin/go-sdk/logger"
	"food_delivery/plugin/tokenprovider"
)

// type RenewTokenStorage interface {
// 	GetSession(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*sessionmodel.Session, error)
// }
type RedisTokenStorage interface {
	WLFindRefreshToken(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*common.RedisToken, error)
	WLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
	BLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
	BLFindTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) bool
	WLDelTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
	WLGetKeys(ctx context.Context, conditions map[string]interface{}) ([]string, error)
	WLDelKeys(ctx context.Context, keys []string) error
}
type renewTokenRepo struct {
	// appCtx        appctx.AppContext
	redisStore    RedisTokenStorage
	tokenProvider tokenprovider.Provider

	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewRenewTokenRepo(redisStore RedisTokenStorage,
	tokenProvider tokenprovider.Provider,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration) *renewTokenRepo {
	return &renewTokenRepo{
		redisStore:    redisStore,
		tokenProvider: tokenProvider,

		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}
func (business *renewTokenRepo) RenewAccessToken(ctx context.Context, data *usermodel.RenewAccessTokenRequest) (*usermodel.Account, error) {
	logger := logger.GetCurrent().GetLogger("module.user.repo.renewtoken")
	refreshPayload, err := business.tokenProvider.Validate(data.Refreshtoken)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	isBlacklist := business.redisStore.BLFindTokens(ctx,
		map[string]interface{}{"id": refreshPayload.UserId(),
			common.KeyRedisRefreshToken: data.Refreshtoken})
	logger.Debugf("isBlacklist %v", isBlacklist)
	if isBlacklist {
		// return nil,
		logger.Debugf("vao trong 1")
		keys, _ := business.redisStore.WLGetKeys(ctx, map[string]interface{}{"id": refreshPayload.UserId()})
		if keys != nil {
			logger.Debugf("vao trong 2")
			_ = business.redisStore.WLDelKeys(ctx, keys)
		}
		return nil, common.NewCusUnauthorizedError(nil, "Token invalid", "ErrTokenInvalid")

	}

	conditions := map[string]interface{}{"id": refreshPayload.UserId(), common.KeyRedisRefreshToken: data.Refreshtoken}
	_, err = business.redisStore.WLFindRefreshToken(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity("RedisToken", err)
	}

	payload := &common.TokenPayload{
		UID:     refreshPayload.UserId(),
		URole:   refreshPayload.Role(),
		TokenID: refreshPayload.ID(),
	}
	accessTokenPayload, err := business.tokenProvider.Generate(payload, business.accessTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	refreshTokenPayload, err := business.tokenProvider.Generate(payload, business.refreshTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// delete tokens whitelist
	_ = business.redisStore.WLDelTokens(ctx,
		map[string]interface{}{"id": refreshPayload.UserId(),
			common.KeyRedisAccessToken:  data.AccessToken,
			common.KeyRedisRefreshToken: data.Refreshtoken})

	// save blacklist token

	err = business.redisStore.BLSaveTokens(ctx,
		map[string]interface{}{"id": refreshPayload.UserId(),
			common.KeyRedisAccessToken:  data.AccessToken,
			common.KeyRedisRefreshToken: data.Refreshtoken})
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	// save whitelist new token

	err = business.redisStore.WLSaveTokens(ctx,
		map[string]interface{}{"id": refreshPayload.UserId(),
			common.KeyRedisAccessToken:  accessTokenPayload,
			common.KeyRedisRefreshToken: refreshTokenPayload})

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(&accessTokenPayload, &refreshTokenPayload)
	return account, nil

}
