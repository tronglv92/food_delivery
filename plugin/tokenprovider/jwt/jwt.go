package jwt

import (
	"flag"
	"fmt"
	"food_delivery/common"
	"food_delivery/plugin/tokenprovider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	prefix string
	secret string
}
type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

func NewTokenJWTProvider(prefix string) *jwtProvider {
	return &jwtProvider{
		prefix: prefix,
	}
}

type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayload{
			UID:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return claims.Payload, nil
}
func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}

func (j *jwtProvider) GetPrefix() string {
	return j.prefix
}

func (j *jwtProvider) Get() interface{} {
	return j
}

func (j *jwtProvider) Name() string {
	return "jwt"
}

func (j *jwtProvider) InitFlags() {
	prefix := j.prefix
	if j.prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&j.secret, prefix+"secret", "dogsupercute", "Secret key for JWT.")
}

func (jwtProvider) Configure() error {
	return nil
}

func (jwtProvider) Run() error {
	return nil
}

func (jwtProvider) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}