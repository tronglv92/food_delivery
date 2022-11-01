package middleware

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/tokenprovider"
	"strings"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}
func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	//"Authorization":"Bearer {token}"
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// 1. Get Token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequiredAuth(sc goservice.ServiceContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)

	return func(c *gin.Context) {

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		// db := appCtx.GetMainDBConnection()
		// store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		ctx, span := trace.StartSpan(c.Request.Context(), "middleware.RequiredAuth")
		user, err := authStore.FindUser(ctx, map[string]interface{}{"id": payload.UserId})
		span.End()

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(
				errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()

	}
}
