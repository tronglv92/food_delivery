package userrepo

import (
	"context"
	"time"

	"food_delivery/common"

	usermodel "food_delivery/module/user/model"

	"food_delivery/plugin/tokenprovider"
)

// type RenewTokenStorage interface {
// 	GetSession(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*sessionmodel.Session, error)
// }
type RedisTokenStorage interface {
	FindRefreshToken(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*common.RedisToken, error)
	SaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
	DelTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
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

	refreshPayload, err := business.tokenProvider.Validate(data.Refreshtoken)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	conditions := map[string]interface{}{"id": refreshPayload.UserId(), common.KeyRedisRefreshToken: data.Refreshtoken}
	_, err = business.redisStore.FindRefreshToken(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity("RedisToken", err)
	}

	payload := &common.TokenPayload{
		UID:     refreshPayload.UserId(),
		URole:   refreshPayload.Role(),
		TokenID: refreshPayload.ID(),
	}
	accessToken, err := business.tokenProvider.Generate(payload, business.accessTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	refreshToken, err := business.tokenProvider.Generate(payload, business.refreshTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	_ = business.redisStore.DelTokens(ctx,
		map[string]interface{}{"id": refreshPayload.UserId(),
			common.KeyRedisAccessToken:  data.AccessToken,
			common.KeyRedisRefreshToken: data.Refreshtoken})

	err = business.redisStore.SaveTokens(ctx,
		map[string]interface{}{"id": refreshPayload.UserId(),
			common.KeyRedisAccessToken:  accessToken,
			common.KeyRedisRefreshToken: refreshToken})

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(&accessToken, &refreshToken)
	return account, nil

}
