package handlers

import (
	"food_delivery/common"
	"food_delivery/middleware"
	"food_delivery/module/restaurant/transport/ginrestaurant"
	"food_delivery/module/restaurantlike/transport/ginrstlike"
	"food_delivery/module/upload/uploadtransport/ginupload"
	userstoragegorm "food_delivery/module/user/storage/gorm"
	userstorageredis "food_delivery/module/user/storage/redis"

	gindevicetoken "food_delivery/module/devicetoken/transport/gin"
	usergin "food_delivery/module/user/transport/ginuser"
	"net/http"

	goservice "food_delivery/plugin/go-sdk"

	"food_delivery/plugin/storage/sdkredis"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	dbConn := sc.MustGet(common.DBMain).(*gorm.DB)
	// userStore := userstoragegorm.NewSQLStore(dbConn)
	userStore := userstorageredis.NewAuthUserCache(
		userstoragegorm.NewSQLStore(dbConn),
		sdkredis.NewRedisCache(sc),
	)
	v1 := router.Group("/v1")
	{
		v1.GET("/admin",
			middleware.RequiredAuth(sc, userStore),
			middleware.RequiredRoles(sc, "admin", "mod"),
			func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": 1})
			})

		v1.POST("/register", usergin.Register(sc))
		v1.POST("/auth", usergin.Login(sc))
		v1.GET("/profile", middleware.RequiredAuth(sc, userStore), usergin.Profile(sc))
		v1.POST("/upload", middleware.RequiredAuth(sc, userStore), ginupload.Upload(sc))
		v1.POST("/put-device-token", middleware.RequiredAuth(sc, userStore), gindevicetoken.PutDeviceToken(sc))
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", middleware.RequiredAuth(sc, userStore), ginrestaurant.CreateRestaurant(sc))

			restaurants.GET("", ginrestaurant.ListRestaurant(sc))
			restaurants.PUT("/:id", ginrestaurant.UpdateRestaurantHandler(sc))
			restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(sc))
			restaurants.GET("/:id", ginrestaurant.GetRestaurant(sc))

			restaurants.POST("/:id/like", middleware.RequiredAuth(sc, userStore), ginrstlike.UserLikeRestaurant(sc))
			restaurants.DELETE("/:id/dislike", middleware.RequiredAuth(sc, userStore), ginrstlike.UserDislikeRestaurant(sc))
			restaurants.GET("/:id/like", middleware.RequiredAuth(sc, userStore), ginrstlike.ListUser(sc))
		}
	}
}
