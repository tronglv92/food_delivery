package ginrstlike

import (
	"food_delivery/common"
	rstlikebiz "food_delivery/module/restaurantlike/biz"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"food_delivery/plugin/pubsub"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// /v1/restaurants/:id/like
func UserLikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
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

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		ps := sc.MustGet(common.PluginRabbitMQ).(pubsub.Pubsub)
		store := restaurantlikestorage.NewSQLStore(db)
		// incStore := restaurantlikestorage.NewSQLStore(db)

		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, ps)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
