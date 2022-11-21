package common

import "log"

type DbType int

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)
const (
	CurrentUser          = "user"
	DBMain               = "mysql"
	DBMongo              = "mongo"
	PluginUserService    = "user-service"
	JWTProvider          = "jwt"
	PluginPubSub         = "pubsub"
	PluginNATS           = "nats"
	PluginRedis          = "redis"
	PluginES             = "elastic-search"
	PluginGrpcServer     = "grpc-server"
	PluginGrpcUserClient = "grpc-user-client"
	PluginAWS            = "aws"
	PluginFCM            = "fcm"

	TopicUserLikeRestaurant    = "restaurant.liked"
	TopicUserDislikeRestaurant = "restaurant.disliked"
	TopicSendNotification      = "fcm.notification"
)

const (
	DBMongoName     = "food_delivery"
	UsersCollection = "Users"
)

// const (
// 	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
// 	TopicUserDislikeRestaurant = "TopicUserDislikeRestaurant"
// )

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error", err)
	}
}

type TokenPayload struct {
	UID   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.UID
}

func (p TokenPayload) Role() string {
	return p.URole
}
