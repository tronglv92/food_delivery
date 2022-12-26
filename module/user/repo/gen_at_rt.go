package userrepo

import (
	"context"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/go-sdk/logger"
	"food_delivery/plugin/tokenprovider"
	"time"
)

type RedisGenStorage interface {
	WLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
}
type SessionGenStorage interface {
	CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error
}
type genToken struct {
	tokenProvider tokenprovider.Provider
	storeRedis    RedisLoginAppleStorage

	storeSession       SessionAppleStorage
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewGenToken(

	tokenProvider tokenprovider.Provider,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,

	storeRedis RedisLoginAppleStorage,
	storeSession SessionAppleStorage,
) *genToken {
	return &genToken{

		tokenProvider:      tokenProvider,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,

		storeRedis:   storeRedis,
		storeSession: storeSession,
	}
}
func (repo *genToken) GenAtRt(
	ctx context.Context,
	payload tokenprovider.TokenPayload,
	userAgent string,
	clientIp string,
) (*usermodel.Account, error) {
	_ = logger.GetCurrent().GetLogger("module.user.biz.gen_at_rt.go")

	accessTokenResult, err := repo.tokenProvider.Generate(payload, repo.accessTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := repo.tokenProvider.Generate(payload, repo.refreshTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(&accessTokenResult, &refreshToken)
	_ = repo.storeRedis.WLSaveTokens(ctx, map[string]interface{}{

		"id":                        payload.UserId(),
		common.KeyRedisAccessToken:  accessTokenResult,
		common.KeyRedisRefreshToken: refreshToken})

	expiresAt := time.Now().UTC().Add(repo.refreshTokenExpiry)
	session := sessionmodel.SessionCreate{
		ID:           payload.ID().String(),
		UserId:       payload.UserId(),
		RefreshToken: refreshToken.GetToken(),
		UserAgent:    userAgent,
		ClientIp:     clientIp,
		IsBlocked:    0,
		ExpiresAt:    &expiresAt,
	}
	repo.storeSession.CreateSession(ctx, &session)
	return account, nil
}
