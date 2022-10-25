package middleware

import (
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(*gin.Context) {

	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}
		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}
		c.Next()
	}
}
