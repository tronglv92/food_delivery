package ginuser

import (
	"food_delivery/common"
	userbiz "food_delivery/module/user/biz"
	usermodel "food_delivery/module/user/model"
	userrepo "food_delivery/module/user/repo"
	userstore "food_delivery/module/user/storage/gorm"
	storeredis "food_delivery/module/user/storage/redis"

	"net/http"

	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/storage/sdkredis"
	"food_delivery/plugin/tokenprovider"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RenewAccessToken(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request usermodel.RenewAccessTokenRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)

		// store := sessionstore.NewSQLStore(db)
		storeRedis := storeredis.NewAuthUserCache(
			userstore.NewSQLStore(db),
			sdkredis.NewRedisCache(sc),
		)

		repo := userrepo.NewRenewTokenRepo(storeRedis, tokenProvider, common.AccessTokenDuration, common.RefreshTokenDuration)
		biz := userbiz.NewRenewTokenBiz(repo)
		account, err := biz.RenewAccessToken(c.Request.Context(), &request)

		if err != nil {
			panic(common.NewUnauthorized(
				err, "unauthorized", err.Error(), "ErrUnauthorized"))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
