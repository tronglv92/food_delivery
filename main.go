package main

import "food_delivery/cmd"

// func main() {

// 	l := logrus.New()
// 	l.SetFormatter(&logrus.JSONFormatter{})
// 	l.SetOutput(os.Stdout)
// 	l.Logln(logrus.InfoLevel, "Hello word")

// 	dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
// 	s3BucketName := "g123456-my-bucket"
// 	s3Region := "ap-southeast-1"
// 	s3APIKey := "AKIAZDLBDNV4DS54J4LG"
// 	s3SecretKey := "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"
// 	s3Domain := "https://d3s5ma63l4xcbq.cloudfront.net"
// 	// PASSWORD SECRET KEY
// 	secretKey := "dogsupercute"

// 	// dsn := os.Getenv("MYSQL_CONN_STRING")
// 	// s3BucketName := os.Getenv("S3BucketName")
// 	// s3Region := os.Getenv("S3Region")
// 	// s3APIKey := os.Getenv("S3APIKey")
// 	// s3SecretKey := os.Getenv("S3SecretKey")
// 	// s3Domain := os.Getenv("S3Domain")
// 	// secretKey := os.Getenv("SYSTEM_SECRET")

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// db.Debug()

// 	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
// 	ps := pblocal.NewPubSub()
// 	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

// 	// setup background
// 	// subscriber.Setup(appContext, context.Background())
// 	_ = subscriber.NewEngine(appContext).Start()

// 	r := gin.Default()

// 	r.StaticFile("/demo/", "./demo.html")

// 	r.Use(middleware.Recover(appContext))

// 	userStore := userstore.NewSQLStore(appContext.GetMainDBConnection())
// 	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)

// 	v1 := r.Group("v1")
// 	mainroute.SetupRoute(appContext, v1, userCachingStore)

// 	admin := v1.Group("/admin", middleware.RequiredAuth(appContext, userCachingStore), middleware.RoleRequired(appContext, "admin", "mod"))
// 	mainroute.SetupAdminRoute(appContext, admin)

// 	//startSocketIOServer(r, appContext)
// 	rtEngine := skio.NewEngine()
// 	appContext.SetRealtimeEngine(rtEngine)

// 	_ = rtEngine.Run(appContext, r)

// 	// Config exporter Jaeger
// 	agentEndpointURI := "localhost:6831"

// 	je, err := jg.NewExporter(jg.Options{
// 		AgentEndpoint: agentEndpointURI,
// 		// CollectorEndpoint: collectorEndpointURI,
// 		// ServiceName:       "demo",
// 		Process: jg.Process{ServiceName: "G05-Food-Delivery"},
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	trace.RegisterExporter(je)
// 	trace.ApplyConfig(trace.Config{
// 		DefaultSampler: trace.ProbabilitySampler(1),
// 	})

// 	// // Setup gRPC Server
// 	// address := "0.0.0.0:50051"
// 	// lis, err := net.Listen("tcp", address)

// 	// if err != nil {
// 	// 	log.Fatalf("Error %v", err)
// 	// }
// 	// fmt.Printf("Server is listening on %v ...", address)

// 	// s := grpc.NewServer()

// 	// gRPCStore := restaurantlikestorage.NewGRPCStore(restaurantlikestorage.NewSQLStore(db))

// 	// demo.RegisterRestaurantLikeServiceServer(s, gRPCStore)

// 	// // if err := s.Server(lis); err != nil {
// 	// // 	log.Fatalln(err)
// 	// // }
// 	// go s.Serve(lis)

// 	// // Setup gRPC Connection
// 	// opts := grpc.WithInsecure()
// 	// cc, err := grpc.Dial("localhost:50051", opts)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// log.Println("Connect to gRPC server successfully")

// 	// defer cc.Close()
// 	// appContext.SetGRPCClientConnection(cc)

// 	// // Setup gRPC Gateway

// 	// // Create a client connection to the gRPC server we just started
// 	// // This is where the gRPC-Gateway proxies the requests
// 	// conn, err := grpc.DialContext(
// 	// 	context.Background(),
// 	// 	"0.0.0.0:50051",
// 	// 	grpc.WithBlock(),
// 	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	// )
// 	// if err != nil {
// 	// 	log.Fatalln("Failed to dial server:", err)
// 	// }

// 	// gwmux := runtime.NewServeMux()
// 	// // Register Greeter
// 	// err = demo.RegisterRestaurantLikeServiceHandler(context.Background(), gwmux, conn)
// 	// if err != nil {
// 	// 	log.Fatalln("Failed to register gateway:", err)
// 	// }

// 	// gwServer := &http.Server{
// 	// 	Addr:    ":8090",
// 	// 	Handler: gwmux,
// 	// }

// 	// log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
// 	// go gwServer.ListenAndServe()

// 	http.ListenAndServe(":8080", &ochttp.Handler{
// 		Handler: r,
// 	},
// 	)
// 	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

func main() {
	cmd.Execute()
	// dsn := os.Getenv("MYSQL_CONN_STRING")
	// fmt.Println("dsn ", dsn)
}
