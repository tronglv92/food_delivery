package usermodel

import "food_delivery/common"

type LoginGoogleRequest struct {
	AccessToken string `json:"access_token" form:"access_token" `
}
type LoginGoogleResponse struct {
	// Data []common.SimpleUser `json:"data"`
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Photo     string `json:"picture"`
}
type UserGoogleCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Role            string `json:"-" gorm:"column:role;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	GoogleId        string `json:"-" gorm:"column:gg_id;"`
	// Avatar          string `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserGoogleCreate) TableName() string {
	return User{}.TableName()
}
