package userrepo

import (
	"context"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
	userBiz "food_delivery/module/user/biz"
	usermodel "food_delivery/module/user/model"

	"food_delivery/plugin/tokenprovider"
	"time"

	"github.com/google/uuid"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type SessionStorage interface {
	CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error
}
type RedisUserStorage interface {
	WLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
}
type loginRepo struct {
	// appCtx        appctx.AppContext
	storeUser          LoginStorage
	storeRedis         RedisUserStorage
	storeSession       SessionStorage
	tokenProvider      tokenprovider.Provider
	hasher             userBiz.Hasher
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewLoginRepo(
	storeUser LoginStorage,
	storeRedis RedisUserStorage,
	storeSession SessionStorage,
	tokenProvider tokenprovider.Provider,
	hasher userBiz.Hasher,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration) *loginRepo {
	return &loginRepo{
		storeUser:          storeUser,
		storeRedis:         storeRedis,
		storeSession:       storeSession,
		tokenProvider:      tokenProvider,
		hasher:             hasher,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1 Access token and refresh token
// 4. Return tokens
func (repo *loginRepo) Login(ctx context.Context, userAgent string, clientIp string, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := repo.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := repo.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &common.TokenPayload{
		UID:     user.Id,
		URole:   user.Role,
		TokenID: tokenID,
	}

	accessToken, err := repo.tokenProvider.Generate(payload, repo.accessTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := repo.tokenProvider.Generate(payload, repo.refreshTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(&accessToken, &refreshToken)

	_ = repo.storeRedis.WLSaveTokens(ctx, map[string]interface{}{

		"id":                        user.Id,
		common.KeyRedisAccessToken:  accessToken,
		common.KeyRedisRefreshToken: refreshToken})

	expiresAt := time.Now().UTC().Add(repo.refreshTokenExpiry)
	session := sessionmodel.SessionCreate{
		ID:           tokenID.String(),
		UserId:       user.Id,
		RefreshToken: refreshToken.GetToken(),
		UserAgent:    userAgent,
		ClientIp:     clientIp,
		IsBlocked:    0,
		ExpiresAt:    &expiresAt,
	}
	repo.storeSession.CreateSession(ctx, &session)
	return account, nil

}
