package cmd

import (
	"fmt"
	handlers "food_delivery/cmd/handler"
	"food_delivery/common"
	"food_delivery/middleware"
	"food_delivery/plugin/sdkgorm"
	"food_delivery/plugin/tokenprovider/jwt"
	"net/http"
	"os"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.JWTProvider)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}
		// appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)
		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())
			engine.GET("/ping", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"data": "pong"})
			})
			handlers.MainRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	// TransAddPoint outenv as a sub command
	rootCmd.AddCommand(outEnvCmd)
	// rootCmd.AddCommand(cronjob)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
