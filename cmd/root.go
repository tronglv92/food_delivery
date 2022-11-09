package cmd

import (
	"fmt"
	handlers "food_delivery/cmd/handler"
	"food_delivery/common"
	"food_delivery/middleware"
	"food_delivery/module/user/store/grpcstore"
	"food_delivery/plugin/appredis"
	appgrpc "food_delivery/plugin/remotecall/grpc"
	"food_delivery/plugin/sdkgorm"
	"food_delivery/plugin/tokenprovider/jwt"
	appnats "food_delivery/pubsub/nats"
	"food_delivery/pubsub/pblocal"
	"net/http"
	"os"

	user "food_delivery/proto"

	userstorage "food_delivery/module/user/store"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.JWTProvider)),
		goservice.WithInitRunnable(pblocal.NewPubSub(common.PluginPubSub)),
		goservice.WithInitRunnable(appnats.NewNATS(common.PluginNATS)),
		goservice.WithInitRunnable(appredis.NewRedisDB("redis", common.PluginRedis)),
		goservice.WithInitRunnable(appgrpc.NewGRPCServer(common.PluginGrpcServer)),
		goservice.WithInitRunnable(appgrpc.NewUserClient(common.PluginGrpcUserClient)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")

		service.MustGet(common.PluginGrpcServer).(interface {
			SetRegisterHdl(hdl func(*grpc.Server))
		}).SetRegisterHdl(func(server *grpc.Server) {
			dbConn := service.MustGet(common.DBMain).(*gorm.DB)
			user.RegisterUserServiceServer(server, grpcstore.NewGRPCStore(userstorage.NewSQLStore(dbConn)))
		})

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
			handlers.InternalRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	// TransAddPoint outenv as a sub command
	rootCmd.AddCommand(outEnvCmd)
	rootCmd.AddCommand(cronjob)
	rootCmd.AddCommand(startSubUserLikedRestaurantCmd)
	rootCmd.AddCommand(startSubUserDislikedRestaurantCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
