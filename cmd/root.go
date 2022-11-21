package cmd

import (
	"fmt"
	handlers "food_delivery/cmd/handler"
	"food_delivery/common"
	"food_delivery/middleware"
	"food_delivery/plugin/aws"
	"food_delivery/plugin/fcm"
	appnats "food_delivery/plugin/pubsub/nats"
	"food_delivery/plugin/pubsub/pblocal"
	appgrpc "food_delivery/plugin/remotecall/grpc"
	"food_delivery/plugin/remotecall/restful"

	userstorage "food_delivery/module/user/storage/gorm"
	usergrpc "food_delivery/module/user/storage/grpc"
	"food_delivery/plugin/jaeger"
	"food_delivery/plugin/storage/sdkes"
	"food_delivery/plugin/storage/sdkgorm"
	"food_delivery/plugin/storage/sdkmgo"
	"food_delivery/plugin/storage/sdkredis"
	"food_delivery/plugin/tokenprovider/jwt"
	user "food_delivery/proto"
	"log"
	"net/http"
	"os"
	"time"

	goservice "food_delivery/plugin/go-sdk"

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
		goservice.WithInitRunnable(sdkmgo.NewMongoDB("mongoDB", common.DBMongo)),
		goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.JWTProvider)),
		goservice.WithInitRunnable(restful.NewUserService()),
		goservice.WithInitRunnable(pblocal.NewPubSub(common.PluginPubSub)),
		goservice.WithInitRunnable(appnats.NewNATS(common.PluginNATS)),
		goservice.WithInitRunnable(sdkredis.NewRedisDB("redis", common.PluginRedis)),
		goservice.WithInitRunnable(sdkes.NewES("elastic", common.PluginES)),
		goservice.WithInitRunnable(jaeger.NewJaeger("g05-Food-Delivery")),
		goservice.WithInitRunnable(aws.New(common.PluginAWS)),
		goservice.WithInitRunnable(fcm.New(common.PluginFCM)),
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
			user.RegisterUserServiceServer(server, usergrpc.NewGRPCStore(userstorage.NewSQLStore(dbConn)))
		})

		initServiceWithRetry(service, 10)

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

func initServiceWithRetry(s goservice.Service, retry int) {
	var err error
	for i := 1; i <= retry; i++ {
		if err = s.Init(); err != nil {
			s.Logger("service").Errorf("error when starting service: %s", err.Error())
			time.Sleep(time.Second * 3)
			continue
		} else {
			break
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
