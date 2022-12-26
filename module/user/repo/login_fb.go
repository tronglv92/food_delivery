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

type LoginFacebookStorage interface {
	LoginFacebook(ctx context.Context, accessToken string) (*usermodel.LoginFacebookResponse, error)
}
type UserFacebookStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUserFacebook(ctx context.Context, data *usermodel.UserFacebookCreate) error
}
type HasherFacebook interface {
	Hash(data string) string
}
type RedisLoginFacebookStorage interface {
	WLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
}
type SessionFacebookStorage interface {
	CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error
}
type loginFacebookRepo struct {
	loginFacebookStorage LoginFacebookStorage
	userStorage          UserFacebookStorage
	tokenProvider        tokenprovider.Provider
	storeRedis           RedisLoginFacebookStorage
	hasher               HasherFacebook
	storeSession         SessionFacebookStorage
	accessTokenExpiry    time.Duration
	refreshTokenExpiry   time.Duration
}

func NewLoginFacebookRepo(
	loginFacebookStorage LoginFacebookStorage,
	userStorage UserFacebookStorage,
	tokenProvider tokenprovider.Provider,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
	hasher HasherPassword,
	storeRedis RedisLoginFacebookStorage,
	storeSession SessionFacebookStorage,
) *loginFacebookRepo {
	return &loginFacebookRepo{
		loginFacebookStorage: loginFacebookStorage,
		userStorage:          userStorage,
		tokenProvider:        tokenProvider,
		accessTokenExpiry:    accessTokenExpiry,
		refreshTokenExpiry:   refreshTokenExpiry,
		hasher:               hasher,
		storeRedis:           storeRedis,
		storeSession:         storeSession,
	}
}
func (repo *loginFacebookRepo) LoginFacebook(ctx context.Context, accessToken string, userAgent string, clientIp string) (*usermodel.Account, error) {
	logger := logger.GetCurrent().GetLogger("module.user.biz.login_fb.go")
	loginResponse, err := repo.loginFacebookStorage.LoginFacebook(ctx, accessToken)
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
		data := &usermodel.UserFacebookCreate{
			Email:      loginResponse.Email,
			Password:   repo.hasher.Hash(fmt.Sprintf(common.PasswordFacebok, loginResponse.Email)),
			FirstName:  loginResponse.Name,
			LastName:   loginResponse.Name,
			Role:       common.RoleUser,
			FacebookId: loginResponse.Id,
		}
		data.Salt = salt
		data.Status = 1
		if err := repo.userStorage.CreateUserFacebook(ctx, data); err != nil {
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
