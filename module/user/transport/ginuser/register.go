package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/hasher"
	userbiz "food_delivery/module/user/biz"
	usermodel "food_delivery/module/user/model"
	userstorage "food_delivery/module/user/storage/gorm"
	userstoragemgo "food_delivery/module/user/storage/mongo"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Register(sc goservice.ServiceContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		client := sc.MustGet(common.DBMongo).(*mongo.Client)
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		_ = userstorage.NewSQLStore(db)
		store := userstoragemgo.NewMongoStore(client)

		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(common.DbTypeUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
