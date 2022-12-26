package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/hasher"
	sessionstore "food_delivery/module/session/storage/gorm"
	usermodel "food_delivery/module/user/model"
	userrepo "food_delivery/module/user/repo"
	userstoregorm "food_delivery/module/user/storage/gorm"
	storeredis "food_delivery/module/user/storage/redis"
	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/loginapple"
	"food_delivery/plugin/storage/sdkredis"
	"food_delivery/plugin/tokenprovider"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginApple(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request usermodel.LoginAppleRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		loginApple := sc.MustGet(common.PluginLoginApple).(loginapple.LoginApple)
		db := sc.MustGet(common.DBMain).(*gorm.DB)

		useGorm := userstoregorm.NewSQLStore(db)
		tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)
		md5 := hasher.NewMd5Hash()
		storeSession := sessionstore.NewSQLStore(db)
		storeRedis := storeredis.NewAuthUserCache(
			useGorm,
			sdkredis.NewRedisCache(sc),
		)
		repo := userrepo.NewLoginAppleRepo(loginApple, useGorm, tokenProvider, common.AccessTokenDuration, common.RefreshTokenDuration, md5, storeRedis, storeSession)
		account, err := repo.LoginApple(c.Request.Context(), request.AccessToken, request.Name, c.Request.UserAgent(), c.ClientIP())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
