package ginrstlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	rstlikebiz "food_delivery/module/restaurantlike/biz"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DELETE /v1/restaurants/:id/dislike
func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}