package ginuser

import (
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
	sessionstore "food_delivery/module/session/storage/gorm"
	userbiz "food_delivery/module/user/biz"

	"net/http"

	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/tokenprovider"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RenewAccessToken(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request sessionmodel.RenewAccessTokenRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)

		store := sessionstore.NewSQLStore(db)

		business := userbiz.NewRenewTokenBusiness(store, tokenProvider, common.AccessTokenDuration, common.RefreshTokenDuration)
		account, err := business.RenewAccessToken(c.Request.Context(), &request)

		if err != nil {
			panic(common.NewUnauthorized(
				err, "unauthorized", err.Error(), "ErrUnauthorized"))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
