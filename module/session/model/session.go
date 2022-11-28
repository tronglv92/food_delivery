package sessionmodel

import (
	"errors"
	"food_delivery/common"
	"strings"
	"time"
)

const EntityName = "Session"

type Session struct {
	ID           string     `json:"id" gorm:"column:id;"`
	Status       int        `json:"status" gorm:"column:status;"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt     *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	UserId       int        `json:"user_id" gorm:"column:user_id;"`
	RefreshToken string     `json:"refresh_token" gorm:"column:refresh_token;"`
	UserAgent    string     `json:"user_agent" gorm:"column:user_agent;"`
	ClientIp     string     `json:"client_ip" gorm:"column:client_ip;"`
	IsBlocked    int        `json:"is_blocked" gorm:"column:is_blocked;"`
	ExpiresAt    time.Time  `json:"expires_at" gorm:"column:expires_at;"`
}

func (u *Session) GetSessionId() string {
	return u.ID
}
func (u *Session) GetUserId() int {
	return u.UserId
}

func (Session) TableName() string {
	return "sessions"
}

type SessionCreate struct {
	ID           string     `json:"id" gorm:"column:id;"`
	Status       int        `json:"status" gorm:"column:status;"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt     *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	UserId       int        `json:"user_id" gorm:"column:user_id;"`
	RefreshToken string     `json:"refresh_token" gorm:"column:refresh_token;"`
	UserAgent    string     `json:"user_agent" gorm:"column:user_agent;"`
	ClientIp     string     `json:"client_ip" gorm:"column:client_ip;"`
	IsBlocked    int        `json:"is_blocked" gorm:"column:is_blocked;"`
	ExpiresAt    *time.Time `json:"expires_at" gorm:"column:expires_at;"`
}

func (SessionCreate) TableName() string {
	return Session{}.TableName()
}
func (s *SessionCreate) PrepareForInsert() {
	now := time.Now().UTC()

	s.Status = 1
	s.CreatedAt = &now
	s.UpdateAt = &now
}
func (data *SessionCreate) Validate() error {
	data.ID = strings.TrimSpace(data.ID)

	data.RefreshToken = strings.TrimSpace(data.RefreshToken)
	data.UserAgent = strings.TrimSpace(data.UserAgent)
	data.ClientIp = strings.TrimSpace(data.ClientIp)

	if data.ID == "" {
		return ErrIdEmpty
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

type RenewAccessTokenRequest struct {
	Refreshtoken string `json:"refresh_token" form:"refresh_token" `
}

func (RenewAccessTokenRequest) TableName() string {
	return Session{}.TableName()
}

var (
	ErrSessionIsBlocked = common.NewCustomError(
		errors.New("session is blocked"),
		"session is blocked",
		"ErrSessionIsBlocked",
	)
	ErrIncorrectSessionUser = common.NewCustomError(
		errors.New("incorrect session user"), "incorrect session user", "ErrIncorrectSessionUser",
	)
	ErrMismatchSessionToken = common.NewCustomError(
		errors.New("mismatched session token"), "mismatched session token", "ErrMismatchSessionToken",
	)
	ErrExpiredSession = common.NewCustomError(
		errors.New("expired session"), "expired session", "ErrExpiredSession",
	)
	ErrIdEmpty = common.NewCustomError(
		errors.New("ID is empty"), "ID is empty", "ErrIdEmpty",
	)
	ErrRefreshTokenEmpty = common.NewCustomError(
		errors.New("Refresh Token is empty"), "Refresh Token is empty", "ErrRefreshTokenEmpty",
	)
	ErrUserAgentEmpty = common.NewCustomError(
		errors.New("User Agent is empty"), "User Agent is empty", "ErrUserAgentEmpty",
	)
	ErrClientIpEmpty = common.NewCustomError(
		errors.New("Client IP is empty"), "Client IP is empty", "ErrClientIpEmpty",
	)
)
