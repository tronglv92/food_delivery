package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/hasher"
	sessionstore "food_delivery/module/session/storage/gorm"
	usermodel "food_delivery/module/user/model"
	userrepo "food_delivery/module/user/repo"
	userstore "food_delivery/module/user/storage/gorm"
	userstoregorm "food_delivery/module/user/storage/gorm"
	storeredis "food_delivery/module/user/storage/redis"
	userstoreremote "food_delivery/module/user/storage/remote/restful"
	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/storage/sdkredis"
	"food_delivery/plugin/tokenprovider"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

func LoginGoogle(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request usermodel.LoginGoogleRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		client := sc.MustGet(common.PluginRestService).(*resty.Client)
		db := sc.MustGet(common.DBMain).(*gorm.DB)

		loginGoogleStorage := userstoreremote.NewUserRestfulStore(client)
		userStorage := userstoregorm.NewSQLStore(db)
		tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)
		md5 := hasher.NewMd5Hash()
		storeSession := sessionstore.NewSQLStore(db)
		storeRedis := storeredis.NewAuthUserCache(
			userstore.NewSQLStore(db),
			sdkredis.NewRedisCache(sc),
		)
		repo := userrepo.NewLoginGoogleRepo(loginGoogleStorage, userStorage, tokenProvider, common.AccessTokenDuration, common.RefreshTokenDuration, md5, storeRedis, storeSession)
		account, err := repo.LoginGoogle(c.Request.Context(), request.AccessToken, c.Request.UserAgent(), c.ClientIP())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
