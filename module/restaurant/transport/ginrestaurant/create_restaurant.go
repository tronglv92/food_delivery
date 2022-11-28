package ginrestaurant

import (
	"context"
	"food_delivery/common"
	restaurantbiz "food_delivery/module/restaurant/biz"
	restaurantmodel "food_delivery/module/restaurant/model"
	restaurantrepo "food_delivery/module/restaurant/repository"
	restaurantstorage "food_delivery/module/restaurant/storage/gorm"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/pubsub"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		// arr := []int{}
		// log.Println(arr[0])
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate
		if err := ctx.ShouldBind(&data); err != nil {

			panic(err)
		}

		data.UserId = requester.GetUserId()

		deviceTokenService := sc.MustGet(common.PluginGrpcDeviceTokenClient).(interface {
			GetDeviceTokens(ctx context.Context, userId int32) ([]common.SimpleDeviceToken, error)
		})
		store := restaurantstorage.NewSQLStore(db)
		ps := sc.MustGet(common.PluginNATS).(pubsub.Pubsub)
		repo := restaurantrepo.NewCreateRestaurantRepo(store, deviceTokenService, ps)
		biz := restaurantbiz.NewCreateRestaurantBiz(repo)
		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(common.DbTypeRestaurant)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
