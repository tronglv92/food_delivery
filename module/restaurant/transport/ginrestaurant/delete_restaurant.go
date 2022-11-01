package ginrestaurant

import (
	"food_delivery/common"
	restaurantbiz "food_delivery/module/restaurant/biz"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		//id, err := strconv.Atoi(ctx.Param("id"))

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)
		if err := biz.DeleteRestaurant(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}
