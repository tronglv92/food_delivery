package ginuser

import (
	"food_delivery/common"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
)

func Profile(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
