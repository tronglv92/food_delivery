package userrepo

import (
	"context"
	"fmt"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/go-sdk/logger"
	"food_delivery/plugin/tokenprovider"
	"time"

	"github.com/google/uuid"
)

type LoginGoogleStorage interface {
	LoginGoogle(ctx context.Context, accessToken string) (*usermodel.LoginGoogleResponse, error)
}
type UserStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUserGoogle(ctx context.Context, data *usermodel.UserGoogleCreate) error
}
type HasherPassword interface {
	Hash(data string) string
}
type RedisLoginGoogleStorage interface {
	WLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
}
type SessionGoogleStorage interface {
	CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error
}
type loginGoogleRepo struct {
	loginGoogleStorage LoginGoogleStorage
	userStorage        UserStorage
	tokenProvider      tokenprovider.Provider
	storeRedis         RedisLoginGoogleStorage
	hasher             HasherPassword
	storeSession       SessionGoogleStorage
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewLoginGoogleRepo(
	loginGoogleStorage LoginGoogleStorage,
	userStorage UserStorage,
	tokenProvider tokenprovider.Provider,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
	hasher HasherPassword,
	storeRedis RedisLoginGoogleStorage,
	storeSession SessionGoogleStorage,
) *loginGoogleRepo {
	return &loginGoogleRepo{
		loginGoogleStorage: loginGoogleStorage,
		userStorage:        userStorage,
		tokenProvider:      tokenProvider,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
		hasher:             hasher,
		storeRedis:         storeRedis,
		storeSession:       storeSession,
	}
}
func (repo *loginGoogleRepo) LoginGoogle(ctx context.Context, accessToken string, userAgent string, clientIp string) (*usermodel.Account, error) {
	logger := logger.GetCurrent().GetLogger("module.user.biz.login_gg.go")
	loginResponse, err := repo.loginGoogleStorage.LoginGoogle(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	user, _ := repo.userStorage.FindUser(ctx, map[string]interface{}{"email": loginResponse.Email})
	var payload *common.TokenPayload

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	logger.Debugf("user: %v", user)
	if user != nil {

		payload = &common.TokenPayload{
			UID:     user.Id,
			URole:   user.Role,
			TokenID: tokenID,
		}

	} else {
		salt := common.GenSalt(50)
		data := &usermodel.UserGoogleCreate{
			Email:     loginResponse.Email,
			Password:  repo.hasher.Hash(fmt.Sprintf(common.PasswordGoogle, loginResponse.Email)),
			FirstName: loginResponse.FirstName,
			LastName:  loginResponse.LastName,
			Role:      common.RoleUser,
			GoogleId:  loginResponse.Id,
		}
		data.Salt = salt
		data.Status = 1
		if err := repo.userStorage.CreateUserGoogle(ctx, data); err != nil {
			return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
		}

		logger.Debugf("id: %v", data.Id)
		payload = &common.TokenPayload{
			UID:     data.Id,
			URole:   data.Role,
			TokenID: tokenID,
		}
	}
	genAtRt := NewGenToken(
		repo.tokenProvider,
		repo.accessTokenExpiry,
		repo.refreshTokenExpiry,
		repo.storeRedis,
		repo.storeSession)
	return genAtRt.GenAtRt(ctx, payload, userAgent, clientIp)

}
