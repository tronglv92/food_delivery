package ginrstlike

import (
	"food_delivery/common"
	rstlikebiz "food_delivery/module/restaurantlike/biz"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUser(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := restaurantlikestorage.NewSQLStore(db)
		biz := rstlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(common.DbTypeUser)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
