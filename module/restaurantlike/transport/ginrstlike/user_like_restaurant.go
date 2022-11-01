package ginrstlike

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
)

// /v1/restaurants/:id/like
func UserLikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid, err := common.FromBase58(c.Param("id"))

		// if err != nil {
		// 	panic(common.ErrInvalidRequest(err))
		// }

		// requester := c.MustGet(common.CurrentUser).(common.Requester)

		// data := restaurantlikemodel.Like{
		// 	RestaurantId: int(uid.GetLocalID()),
		// 	UserId:       requester.GetUserId(),
		// }

		// db := sc.MustGet(common.DBMain).(*gorm.DB)
		// store := restaurantlikestorage.NewSQLStore(db)
		// // incStore := restaurantlikestorage.NewSQLStore(db)

		// biz := rstlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		// if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
		// 	panic(err)
		// }
		// c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
