package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)
const (
	CurrentUser = "user"
	DBMain      = "mysql"
	JWTProvider = "jwt"
)

const (
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDislikeRestaurant = "TopicUserDislikeRestaurant"
)

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

func (p *TokenPayload) UserId() int {
	return p.UID
}

func (p *TokenPayload) Role() string {
	return p.URole
}
