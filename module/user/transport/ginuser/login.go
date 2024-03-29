package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/hasher"
	sessionstore "food_delivery/module/session/storage/gorm"
	userbiz "food_delivery/module/user/biz"
	usermodel "food_delivery/module/user/model"
	userrepo "food_delivery/module/user/repo"
	userstore "food_delivery/module/user/storage/gorm"
	storeredis "food_delivery/module/user/storage/redis"
	"food_delivery/plugin/storage/sdkredis"
	"food_delivery/plugin/tokenprovider"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)

		storeUser := userstore.NewSQLStore(db)
		storeSession := sessionstore.NewSQLStore(db)
		storeRedis := storeredis.NewAuthUserCache(
			userstore.NewSQLStore(db),
			sdkredis.NewRedisCache(sc),
		)
		md5 := hasher.NewMd5Hash()

		repo := userrepo.NewLoginRepo(storeUser, storeRedis,storeSession, tokenProvider, md5, common.AccessTokenDuration, common.RefreshTokenDuration)
		biz := userbiz.NewLoginBiz(repo)
		account, err := biz.Login(c.Request.Context(), c.Request.UserAgent(), c.ClientIP(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
