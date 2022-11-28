package ginrestaurant

import (
	"context"
	"food_delivery/common"
	restaurantbiz "food_delivery/module/restaurant/biz"
	restaurantmodel "food_delivery/module/restaurant/model"
	restaurantrepo "food_delivery/module/restaurant/repository"
	resstorageES "food_delivery/module/restaurant/storage/elastic"
	restaurantstorage "food_delivery/module/restaurant/storage/gorm"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

func ListRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		esStore := restaurantstorage.NewSQLStore(db)

		client := sc.MustGet(common.PluginES).(*elastic.Client)
		_ = resstorageES.NewESStore(client)

		// likeStore := restaurantlikestorage.NewSQLStore(db)

		// likeStore := grpcstore.NewGRPCClient(demo.NewRestaurantLikeServiceClient(appCtx.GetGRPCClientConnection()))
		//userStore := userstorage.NewSQLStore(db)
		//userStore := restaurantapi.NewUserApi("http://localhost:4001")

		userService := sc.MustGet(common.PluginGrpcUserClient).(interface {
			GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
		})

		repo := restaurantrepo.NewListRestaurantRepo(esStore, userService)
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
