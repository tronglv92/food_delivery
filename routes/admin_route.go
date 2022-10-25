package mainroute

import (
	"food_delivery/component/appctx"
	"food_delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoute(appContext appctx.AppContext, admin *gin.RouterGroup) {
	admin.GET("/profile", ginuser.Profile(appContext))
}
