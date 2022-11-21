package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/hasher"
	userbiz "food_delivery/module/user/biz"
	usermodel "food_delivery/module/user/model"
	userstore "food_delivery/module/user/storage/gorm"
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

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
