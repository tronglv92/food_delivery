package usertokenmodel

import (
	"errors"
	"food_delivery/common"
	"strings"
	"time"
)

const EntityName = "UserToken"

type UserToken struct {
	Id           int        `json:"id" gorm:"column:id;"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt     *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	UserId       int        `json:"user_id" gorm:"column:user_id;"`
	RefreshToken string     `json:"refresh_token" gorm:"column:refresh_token;"`
	AccessToken  string     `json:"access_token" gorm:"column:access_token;"`
	UserAgent    string     `json:"user_agent" gorm:"column:user_agent;"`
	ClientIp     string     `json:"client_ip" gorm:"column:client_ip;"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}

type UserTokenCreate struct {
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt     *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	UserId       int        `json:"user_id" gorm:"column:user_id;"`
	RefreshToken string     `json:"refresh_token" gorm:"column:refresh_token;"`
	AccessToken  string     `json:"access_token" gorm:"column:access_token;"`
	UserAgent    string     `json:"user_agent" gorm:"column:user_agent;"`
	ClientIp     string     `json:"client_ip" gorm:"column:client_ip;"`
}

func (UserTokenCreate) TableName() string {
	return UserToken{}.TableName()
}

func (s *UserTokenCreate) PrepareForInsert() {
	now := time.Now().UTC()

	s.CreatedAt = &now
	s.UpdateAt = &now
}
func (data *UserTokenCreate) Validate() error {

	data.AccessToken = strings.TrimSpace(data.AccessToken)
	data.RefreshToken = strings.TrimSpace(data.RefreshToken)
	data.UserAgent = strings.TrimSpace(data.UserAgent)
	data.ClientIp = strings.TrimSpace(data.ClientIp)

	if data.UserId == 0 {
		return ErrIdEmpty
	}
	if data.AccessToken == "" {
		return ErrRefreshTokenEmpty
	}
	if data.RefreshToken == "" {
		return ErrRefreshTokenEmpty
	}
	if data.UserAgent == "" {
		return ErrUserAgentEmpty
	}
	if data.ClientIp == "" {
		return ErrClientIpEmpty
	}
	return nil
}

var (
	ErrIdEmpty = common.NewCustomError(
		errors.New("ID is empty"), "ID is empty", "ErrIdEmpty",
	)
	ErrAccessTokenEmpty = common.NewCustomError(
		errors.New("access Token is empty"), "Access Token is empty", "ErrAccessTokenEmpty",
	)
	ErrRefreshTokenEmpty = common.NewCustomError(
		errors.New("refresh Token is empty"), "Refresh Token is empty", "ErrRefreshTokenEmpty",
	)
	ErrUserAgentEmpty = common.NewCustomError(
		errors.New("user Agent is empty"), "User Agent is empty", "ErrUserAgentEmpty",
	)
	ErrClientIpEmpty = common.NewCustomError(
		errors.New("client IP is empty"), "Client IP is empty", "ErrClientIpEmpty",
	)
)
