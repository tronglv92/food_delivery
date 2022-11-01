package middleware

import (
	"food_delivery/common"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
)

func RequiredRoles(sc goservice.ServiceContext, roles ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for i := range roles {
			if u.GetRole() == roles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}
}
