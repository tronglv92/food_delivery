package ginrstlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	rstlikebiz "food_delivery/module/restaurantlike/biz"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /v1/restaurants/:id/like
func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		// incStore := restaurantlikestorage.NewSQLStore(db)

		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
