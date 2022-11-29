package common

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type DbType int

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)
const (
	CurrentUser                 = "user"
	DBMain                      = "mysql"
	DBMongo                     = "mongo"
	PluginUserService           = "user-service"
	JWTProvider                 = "jwt"
	PluginPubSub                = "pubsub"
	PluginNATS                  = "nats"
	PluginRedis                 = "redis"
	PluginES                    = "elastic-search"
	PluginGrpcServer            = "grpc-server"
	PluginGrpcUserClient        = "grpc-user-client"
	PluginGrpcDeviceTokenClient = "grpc-devicetoken-client"
	PluginAWS                   = "aws"
	PluginFCM                   = "fcm"
	PluginRabbitMQ              = "rabbitmq"

	TopicUserLikeRestaurant    = "restaurant.liked"
	TopicUserDislikeRestaurant = "restaurant.disliked"
	TopicSendNotification      = "fcm.notification"
)

const (
	DBMongoName     = "food_delivery"
	UsersCollection = "Users"
)

const (
	AccessTokenDuration  = 1 * time.Hour   // 1 h
	RefreshTokenDuration = 3 * time.Minute // 30 days
	KeyRedisAccessToken  = "access_token"
	KeyRedisRefreshToken = "refresh_token"
	CacheKey             = "user:%d"
	CacheWLKeyAT         = "wl_user:%d:at:%v"
	CacheWLKeyRT         = "wl_user:%d:rt:%v"
	CacheWLPrefixAT      = "wl_user:%d:*"
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
	UID     int       `json:"user_id"`
	URole   string    `json:"role"`
	TokenID uuid.UUID `json:"id"`
}

func (p TokenPayload) UserId() int {
	return p.UID
}

func (p TokenPayload) Role() string {
	return p.URole
}
func (p TokenPayload) ID() uuid.UUID {
	return p.TokenID
}
