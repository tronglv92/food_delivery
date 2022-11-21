package gindevicetoken

import (
	"food_delivery/common"
	devicetokenbiz "food_delivery/module/devicetoken/biz"
	devicetokenmodel "food_delivery/module/devicetoken/model"
	devicetokenstorage "food_delivery/module/devicetoken/storage/gorm"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func PutDeviceToken(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		// arr := []int{}
		// log.Println(arr[0])
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data devicetokenmodel.UserDeviceTokenUpdate
		if err := ctx.ShouldBind(&data); err != nil {

			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := devicetokenstorage.NewSQLStore(db)
		biz := devicetokenbiz.NewPutDeviceTokenBiz(store)
		if err := biz.CreateDeviceToken(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("Update device token success"))
	}
}
