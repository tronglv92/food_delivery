package main

import (
	"food_delivery/component/appctx"
	"food_delivery/component/uploadprovider"
	"food_delivery/middleware"
	"food_delivery/pubsub/pblocal"
	"food_delivery/skio"
	"log"

	mainroute "food_delivery/routes"
	"food_delivery/subscriber"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	s3BucketName := "g012345-food-delivery"
	s3Region := "ap-southeast-1"
	s3APIKey := ""
	s3SecretKey := ""
	s3Domain := "https://d3s5ma63l4xcbq.cloudfront.net"
	// PASSWORD SECRET KEY
	secretKey := "dogsupercute"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	ps := pblocal.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	// setup background
	// subscriber.Setup(appContext, context.Background())
	_ = subscriber.NewEngine(appContext).Start()

	r := gin.Default()

	r.StaticFile("/demo/", "./demo.html")

	r.Use(middleware.Recover(appContext))

	v1 := r.Group("v1")
	mainroute.SetupRoute(appContext, v1)

	admin := v1.Group("/admin", middleware.RequiredAuth(appContext), middleware.RoleRequired(appContext, "admin", "mod"))
	mainroute.SetupAdminRoute(appContext, admin)

	//startSocketIOServer(r, appContext)
	rtEngine := skio.NewEngine()
	appContext.SetRealtimeEngine(rtEngine)

	_ = rtEngine.Run(appContext, r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
// 	server, _ := socketio.NewServer(&engineio.Options{
// 		Transports: []transport.Transport{websocket.Default},
// 	})

// 	server.OnConnect("/", func(s socketio.Conn) error {
// 		// s.SetContext("")
// 		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
// 		s.Join("Shipper")
// 		return nil
// 	})

// 	go func() {
// 		for range time.NewTicker(time.Second).C {
// 			server.BroadcastToRoom("/", "Shipper", "test", "Ahihi")
// 		}
// 	}()

// 	server.OnError("/", func(s socketio.Conn, e error) {
// 		fmt.Println("meet error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 		fmt.Println("closed", reason)
// 	})
// 	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
// 		db := appCtx.GetMainDBConnection()

// 		store := userstore.NewSQLStore(db)

// 		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

// 		payload, err := tokenProvider.Validate(token)

// 		if err != nil {
// 			s.Emit("authentication_failed", err.Error())
// 			s.Close()
// 			return
// 		}

// 		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

// 		if err != nil {
// 			s.Emit("authentication_failed", err.Error())
// 			s.Close()
// 			return
// 		}

// 		if user.Status == 0 {
// 			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
// 			s.Close()
// 			return
// 		}

// 		user.Mask(false)

// 		s.Emit("authenticated", user)

// 	})

// 	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
// 		log.Println("test: ", msg)
// 	})

// 	type Person struct {
// 		Name string `json:"name"`
// 		Age  int    `json:"age"`
// 	}
// 	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
// 		fmt.Println("server receive notice:", p.Name, p.Age)

// 		p.Age = 33
// 		s.Emit("notice", p)
// 	})
// 	// go func() {
// 	// 	if err := server.Serve(); err != nil {
// 	// 		log.Fatalf("socketio listen error: %s\n", err)
// 	// 	}
// 	// }()

// 	go server.Serve()
// 	// defer server.Close()

// 	engine.GET("/socket.io/*any", gin.WrapH(server))
// 	engine.POST("/socket.io/*any", gin.WrapH(server))
// 	// engine.StaticFS("/public", http.Dir("../asset"))

// 	// if err := engine.Run(":8000"); err != nil {
// 	// 	log.Fatal("failed run app: ", err)
// 	// }
// }
