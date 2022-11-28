package userbiz

import (
	"context"
	"time"

	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"

	"food_delivery/plugin/tokenprovider"
)

type RenewTokenStorage interface {
	GetSession(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*sessionmodel.Session, error)
}

type renewTokenBusiness struct {
	// appCtx        appctx.AppContext
	storeUser     RenewTokenStorage
	tokenProvider tokenprovider.Provider

	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}
type renewAccessTokenResponse struct {
	AccessToken *tokenprovider.Token
}

func NewRenewTokenBusiness(storeUser RenewTokenStorage,
	tokenProvider tokenprovider.Provider,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration) *renewTokenBusiness {
	return &renewTokenBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,

		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}
func (business *renewTokenBusiness) RenewAccessToken(ctx context.Context, data *sessionmodel.RenewAccessTokenRequest) (*renewAccessTokenResponse, error) {

	refreshPayload, err := business.tokenProvider.Validate(data.Refreshtoken)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	session, err := business.storeUser.GetSession(ctx, map[string]interface{}{"id": refreshPayload.ID().String()})
	if err != nil {
		return nil, common.ErrCannotGetEntity(sessionmodel.EntityName, err)
	}

	if session.IsBlocked == 1 {
		return nil, sessionmodel.ErrSessionIsBlocked

	}
	if session.UserId != refreshPayload.UserId() {
		return nil, sessionmodel.ErrIncorrectSessionUser
	}
	if session.RefreshToken != data.Refreshtoken {
		return nil, sessionmodel.ErrMismatchSessionToken
	}
	if time.Now().After(session.ExpiresAt) {
		return nil, sessionmodel.ErrExpiredSession
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
	// refreshToken, err := business.tokenProvider.Generate(payload, business.refreshTokenExpiry)
	// if err != nil {
	// 	return nil, common.ErrInternal(err)
	// }
	account := renewAccessTokenResponse{
		AccessToken: &accessToken,
	}
	return &account, nil

}
