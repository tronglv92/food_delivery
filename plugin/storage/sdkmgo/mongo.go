/*
 * @author          Viet Tran <viettranx@gmail.com>
 * @copyright       2019 Viet Tran <viettranx@gmail.com>
 * @license         Apache-2.0
 */

package sdkmgo

import (
	"context"
	"flag"
	"food_delivery/plugin/go-sdk/logger"
	"math"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	defaultDBName  = "defaultMongoDB"
	DefaultMongoDB = getDefaultMongoDB()
)

const retryCount = 10

type MongoDBOpt struct {
	MgoUri       string
	Username     string
	Password     string
	Prefix       string
	Database     string
	PingInterval int // in seconds
}

type mongoDB struct {
	name      string
	logger    logger.Logger
	client    *mongo.Client
	database  *mongo.Database
	isRunning bool
	once      *sync.Once
	*MongoDBOpt
}

func getDefaultMongoDB() *mongoDB {
	return NewMongoDB(defaultDBName, "")
}

func NewMongoDB(name, prefix string) *mongoDB {
	return &mongoDB{
		MongoDBOpt: &MongoDBOpt{
			Prefix: prefix,
		},
		name:      name,
		isRunning: false,
		once:      new(sync.Once),
	}
}

func (mgDB *mongoDB) GetPrefix() string {
	return mgDB.Prefix
}

func (mgDB *mongoDB) Name() string {
	return mgDB.name
}

func (mgDB *mongoDB) InitFlags() {
	prefix := mgDB.Prefix
	if mgDB.Prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&mgDB.MgoUri, prefix+"mgo-uri", "", "MongoDB connection-string. Ex: mongodb://...")

	flag.StringVar(&mgDB.Username, prefix+"mgo_username", "", "MongoDB username. ")
	flag.StringVar(&mgDB.Password, prefix+"mgo_password", "", "MongoDB password. ")
	flag.IntVar(&mgDB.PingInterval, prefix+"mgo-ping-interval", 5, "MongoDB ping check interval")
	flag.Parse()
}

func (mgDB *mongoDB) isDisabled() bool {
	return mgDB.MgoUri == ""
}

func (mgDB *mongoDB) Configure() error {
	if mgDB.isDisabled() || mgDB.isRunning {
		return nil
	}

	mgDB.logger = logger.GetCurrent().GetLogger(mgDB.name)
	mgDB.logger.Info("Connect to Mongodb at ", mgDB.MgoUri, " ...")

	var err error
	mgDB.client, err = mgDB.getConnWithRetry(retryCount)
	if err != nil {
		mgDB.logger.Error("Error connect to mongodb at ", mgDB.MgoUri, ". ", err.Error())
		return err
	}
	mgDB.isRunning = true
	return nil
}

func (mgDB *mongoDB) Cleanup() {
	if mgDB.isDisabled() {
		return
	}

	if mgDB.client != nil {
		mgDB.client.Disconnect(context.TODO())
	}
}

func (mgDB *mongoDB) Run() error {
	return mgDB.Configure()
}

func (mgDB *mongoDB) Stop() <-chan bool {
	if mgDB.client != nil {
		mgDB.client.Disconnect(context.TODO())
	}
	mgDB.isRunning = false

	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (mgDB *mongoDB) Get() interface{} {
	mgDB.once.Do(func() {
		if !mgDB.isRunning && !mgDB.isDisabled() {
			if db, err := mgDB.getConnWithRetry(math.MaxInt32); err == nil {
				mgDB.client = db
				mgDB.isRunning = true
			} else {
				mgDB.logger.Fatalf("%s connection cannot reconnect\n", mgDB.name)
			}
		}
	})

	if mgDB.client == nil {
		return nil
	}
	return mgDB.client
}

func (mgDB *mongoDB) getConnWithRetry(retryCount int) (*mongo.Client, error) {

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	var clientOptions *options.ClientOptions
	if mgDB.Username != "" && mgDB.Password != "" {
		credential := options.Credential{
			Username: mgDB.Username,
			Password: mgDB.Password,
		}
		clientOptions = options.Client().ApplyURI(mgDB.MgoUri).SetAuth(credential)
	} else {
		clientOptions = options.Client().ApplyURI(mgDB.MgoUri)

	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		for {
			time.Sleep(time.Second * 1)
			mgDB.logger.Errorf("Retry to connect %s.\n", mgDB.name)
			client, err = mongo.Connect(context.TODO(), clientOptions)

			if err == nil {
				go mgDB.reconnectIfNeeded()
				break
			}
		}
	} else {
		go mgDB.reconnectIfNeeded()
	}

	return client, err
}

func (mgDB *mongoDB) reconnectIfNeeded() {
	conn := mgDB.client
	for {
		if err := conn.Ping(context.TODO(), readpref.Primary()); err != nil {
			conn.Disconnect(context.TODO())
			mgDB.logger.Errorf("%s connection is gone, try to reconnect\n", mgDB.name)
			mgDB.isRunning = false
			mgDB.once = new(sync.Once)

			// mgDB.Get().(*mgo.Session).Close()
			return
		}
		time.Sleep(time.Second * time.Duration(mgDB.PingInterval))
	}
}
