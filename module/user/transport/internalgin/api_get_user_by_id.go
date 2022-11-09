package userinternal

import (
	"food_delivery/common"
	"net/http"

	userstorage "food_delivery/module/user/store"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserById(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			Ids []int `json:"ids"`
		}

		if err := c.ShouldBind(&param); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := userstorage.NewSQLStore(db)

		result, err := store.GetUsers(c.Request.Context(), param.Ids)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(common.DbTypeUser)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
