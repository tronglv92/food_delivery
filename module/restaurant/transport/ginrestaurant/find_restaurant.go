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

func GetRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant_id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		storage := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(storage)

		data, err := biz.FindRestaurantById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
