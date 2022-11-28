package ginuser

import (
	"food_delivery/common"
	userbiz "food_delivery/module/user/biz"
	userstore "food_delivery/module/user/storage/gorm"
	storeredis "food_delivery/module/user/storage/redis"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/storage/sdkredis"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RemoveRedisToken(sc goservice.ServiceContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		// client := sc.MustGet(common.DBMongo).(*mongo.Client)
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// _ = userstorage.NewSQLStore(db)
		storeRedis := storeredis.NewAuthUserCache(
			userstore.NewSQLStore(db),
			sdkredis.NewRedisCache(sc),
		)

		// md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRemoveTokenBusiness(storeRedis)

		if err := biz.RemoveRedisToken(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("remove tokens of user success"))
	}
}
