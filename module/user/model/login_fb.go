package usermodel

import "food_delivery/common"

type LoginFacebookRequest struct {
	AccessToken string `json:"access_token" form:"access_token" `
}
type DataPicture struct {
	Url string `json:"url"`
}
type Picture struct {
	Data DataPicture `json:"data"`
}
type LoginFacebookResponse struct {
	// Data []common.SimpleUser `json:"data"`
	Id     string  `json:"id"`
	Email  string  `json:"email"`
	Name   string  `json:"name"`
	Avatar Picture `json:"picture"`
}
type UserFacebookCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Role            string `json:"-" gorm:"column:role;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	FacebookId      string `json:"-" gorm:"column:fb_id;"`
	// Avatar          string `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserFacebookCreate) TableName() string {
	return User{}.TableName()
}
