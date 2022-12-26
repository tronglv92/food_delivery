package usermodel

import "food_delivery/common"

type LoginAppleRequest struct {
	AccessToken string `json:"access_token" form:"access_token" `
	Name        string `json:"name" form:"name" `
}

type LoginAppleResponse struct {
	// Data []common.SimpleUser `json:"data"`
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
type UserAppleCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Role            string `json:"-" gorm:"column:role;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	AppleId         string `json:"-" gorm:"column:apple_id;"`
	// Avatar          string `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserAppleCreate) TableName() string {
	return User{}.TableName()
}
