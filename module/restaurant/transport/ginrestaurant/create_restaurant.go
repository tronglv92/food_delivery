package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbiz "food_delivery/module/restaurant/biz"
	restaurantmodel "food_delivery/module/restaurant/model"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appctx.GetMainDBConnection()
		// arr := []int{}
		// log.Println(arr[0])
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate
		if err := ctx.ShouldBind(&data); err != nil {

			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
