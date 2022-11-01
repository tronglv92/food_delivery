package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/hasher"
	userbiz "food_delivery/module/user/biz"
	usermodel "food_delivery/module/user/model"
	userstorage "food_delivery/module/user/store"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Register(sc goservice.ServiceContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
