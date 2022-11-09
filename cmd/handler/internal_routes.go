package handlers

import (
	userinternal "food_delivery/module/user/transport/internalgin"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
)

func InternalRoute(router *gin.Engine, sc goservice.ServiceContext) {
	internal := router.Group("/internal")
	{
		internal.POST("/get-users-by-ids", userinternal.GetUserById(sc))
	}
}
