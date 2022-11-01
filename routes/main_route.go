package mainroute

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup, userStore middleware.AuthenStore) {

	// v1.POST("/upload", ginupload.Upload(appContext))
	// v1.POST("/register", ginuser.Register(appContext))
	// v1.POST("/authenticate", ginuser.Login(appContext))
	// v1.GET("/profile", middleware.RequiredAuth(appContext, userStore), ginuser.Profile(appContext))
	// restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext, userStore))
	// restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	// restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	// restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	// restaurants.POST("/:id/like", ginrstlike.UserLikeRestaurant(appContext))
	// restaurants.DELETE("/:id/dislike", ginrstlike.UserDislikeRestaurant(appContext))
	// restaurants.GET("/:id/like", ginrstlike.ListUser(appContext))
}
