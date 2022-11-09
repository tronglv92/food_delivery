package ginrstlike

import (
	"food_delivery/common"
	rstlikebiz "food_delivery/module/restaurantlike/biz"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"food_delivery/pubsub"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DELETE /v1/restaurants/:id/dislike
func UserDislikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		ps := sc.MustGet(common.PluginNATS).(pubsub.Pubsub)
		store := restaurantlikestorage.NewSQLStore(db)
		// decStore := restaurantstorage.NewSQLStore(db)
		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, ps)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
