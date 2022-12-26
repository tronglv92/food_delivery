package tokenprovider

import (
	"errors"
	"food_delivery/common"

	"time"

	"github.com/google/uuid"
)

type Provider interface {
	Generate(data TokenPayload, expiry time.Duration) (Token, error)
	Validate(token string) (TokenPayload, error)
}

type TokenPayload interface {
	Role() string
	UserId() int
	ID() uuid.UUID
}

type Token interface {
	GetToken() string
	GetExpire() int
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
	ErrExpiredToken = common.NewCustomError(
		errors.New("token has expired"),
		"token has expired",
		"ErrExpiredToken",
	)
)
