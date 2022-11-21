package devicetokenmodel

import (
	"errors"
	"food_delivery/common"
	"strings"
	"time"
)

const EntityName = "UserDeviceToken"

type UserDeviceToken struct {
	common.SQLModel
	Token        string             `json:"token" gorm:"column:token;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	DeviceId     string             `json:"device_id" gorm:"column:device_id;"`
	IsProduction int                `json:"is_production" gorm:"column:is_production;default:0"`
	OS           string             `json:"os" gorm:"column:os;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (UserDeviceToken) TableName() string { return "user_device_tokens" }

type UserDeviceTokenUpdate struct {
	Token        string `json:"token" gorm:"column:token;"`
	IsProduction int    `json:"is_production" gorm:"column:is_production;default:0"`
	OS           string `json:"os" gorm:"column:os;"`
	DeviceId     string `json:"device_id" gorm:"column:device_id;"`
	UserId       int    `json:"-" gorm:"column:user_id"`
}

func (UserDeviceTokenUpdate) TableName() string { return UserDeviceToken{}.TableName() }
func (data *UserDeviceTokenUpdate) Validate() error {
	data.Token = strings.TrimSpace(data.Token)
	data.OS = strings.TrimSpace(data.OS)
	data.DeviceId = strings.TrimSpace(data.DeviceId)
	if data.Token == "" {
		return ErrTokenIsEmpty
	}
	if data.OS == "" {
		return ErrOSIsEmpty
	}
	if data.DeviceId == "" {
		return ErrDeviceIsEmpty
	}
	return nil
}

var (
	ErrTokenIsEmpty = common.NewCustomError(
		errors.New("Token is empty"),
		"Token is empty",
		"ErrTokenIsEmpty",
	)
	ErrDeviceIsEmpty = common.NewCustomError(
		errors.New("DeviceId is empty"),
		"DeviceId is empty",
		"ErrDeviceIdIsEmpty",
	)
	ErrOSIsEmpty = common.NewCustomError(
		errors.New("OS is empty"),
		"OS is empty",
		"ErrOSIsEmpty",
	)
)
