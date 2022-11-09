package ginrestaurant

import (
	"food_delivery/common"
	restaurantbiz "food_delivery/module/restaurant/biz"
	restaurantmodel "food_delivery/module/restaurant/model"
	restaurantrepo "food_delivery/module/restaurant/repository"
	restaurantstorage "food_delivery/module/restaurant/storage"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)

		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)

		// likeStore := grpcstore.NewGRPCClient(demo.NewRestaurantLikeServiceClient(appCtx.GetGRPCClientConnection()))

		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)
		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
