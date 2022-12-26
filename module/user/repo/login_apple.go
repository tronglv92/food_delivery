package userrepo

import (
	"context"
	"fmt"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/go-sdk/logger"
	"food_delivery/plugin/loginapple"
	"food_delivery/plugin/tokenprovider"
	"time"

	"github.com/google/uuid"
)

type UserGorm interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUserApple(ctx context.Context, data *usermodel.UserAppleCreate) error
}
type HasherApple interface {
	Hash(data string) string
}
type RedisLoginAppleStorage interface {
	WLSaveTokens(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) error
}
type SessionAppleStorage interface {
	CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error
}
type loginAppleRepo struct {
	loginApple         loginapple.LoginApple
	userGorm           UserGorm
	tokenProvider      tokenprovider.Provider
	storeRedis         RedisLoginAppleStorage
	hasher             HasherApple
	storeSession       SessionAppleStorage
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewLoginAppleRepo(
	loginApple loginapple.LoginApple,
	userGorm UserGorm,
	tokenProvider tokenprovider.Provider,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
	hasher HasherPassword,
	storeRedis RedisLoginAppleStorage,
	storeSession SessionAppleStorage,
) *loginAppleRepo {
	return &loginAppleRepo{
		loginApple:         loginApple,
		userGorm:           userGorm,
		tokenProvider:      tokenProvider,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
		hasher:             hasher,
		storeRedis:         storeRedis,
		storeSession:       storeSession,
	}
}
func (repo *loginAppleRepo) LoginApple(
	ctx context.Context,
	accessToken string,
	name string,
	userAgent string,
	clientIp string,
) (*usermodel.Account, error) {
	logger := logger.GetCurrent().GetLogger("module.user.biz.login_apple.go")

	loginResponse, err := repo.loginApple.LoginAppleByAccessToken(ctx, accessToken, name)
	if err != nil {

		return nil, err
	}

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// _, span = trace.StartSpan(ctx, "user.find_user_by_apple_id")
	// user, _ := repo.userStorage.FindUser(ctx, map[string]interface{}{"apple_id": loginResponse.Id})
	// span.End()

	var payload *common.TokenPayload

	user, _ := repo.userGorm.FindUser(ctx, map[string]interface{}{"email": loginResponse.Email})

	if user != nil {
		payload = &common.TokenPayload{
			UID:     user.Id,
			URole:   user.Role,
			TokenID: tokenID,
		}

	} else {
		salt := common.GenSalt(50)
		data := &usermodel.UserAppleCreate{
			Email:     loginResponse.Email,
			Password:  repo.hasher.Hash(fmt.Sprintf(common.PasswordApple, loginResponse.Email)),
			FirstName: loginResponse.Name,
			LastName:  loginResponse.Name,
			Role:      common.RoleUser,
			AppleId:   loginResponse.Id,
		}
		data.Salt = salt
		data.Status = 1

		if err := repo.userGorm.CreateUserApple(ctx, data); err != nil {

			return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
		}

		payload = &common.TokenPayload{
			UID:     data.Id,
			URole:   data.Role,
			TokenID: tokenID,
		}
	}
	logger.Debugf("user: %v", user)

	genAtRt := NewGenToken(
		repo.tokenProvider,
		repo.accessTokenExpiry,
		repo.refreshTokenExpiry,
		repo.storeRedis,
		repo.storeSession)

	return genAtRt.GenAtRt(ctx, payload, userAgent, clientIp)
}
