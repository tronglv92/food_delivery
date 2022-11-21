package restaurantmodel

import (
	"errors"
	"food_delivery/common"
	"strings"
)

const EntityName = "restaurants"

type Restaurant struct {
	common.SQLModel
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`

	UserId      int                `json:"-" gorm:"column:user_id"`
	User        *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikedCount  int                `json:"liked_count" gorm:"column:liked_count;"`
	FakeOwnerId *common.UID        `json:"user_id" gorm:"-"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.SQLModel.Mask(common.DbTypeRestaurant)

	fakeOwnerId := common.NewUID(uint32(r.UserId), int(common.DbTypeUser), 1)
	r.FakeOwnerId = &fakeOwnerId

	if v := r.User; v != nil {
		v.Mask(common.DbTypeUser)
	}
}

type RestaurantES struct {
	Name       string `json:"name"`
	Addr       string `json:"addr"`
	UserId     int    `json:"user_id"`
	Id         int    `json:"id"`
	LikedCount int    `json:"liked_count"`
}

func (r *RestaurantES) ToRestaurant() Restaurant {

	uid := common.NewUID(uint32(r.Id), int(common.DbTypeRestaurant), 1)

	fakeOwnerId := common.NewUID(uint32(r.UserId), int(common.DbTypeRestaurant), 1)
	rs := Restaurant{
		SQLModel: common.SQLModel{
			Id:     int(uid.GetLocalID()),
			FakeId: &uid,
		},
		Name:        r.Name,
		Addr:        r.Addr,
		UserId:      r.UserId,
		LikedCount:  r.LikedCount,
		FakeOwnerId: &fakeOwnerId,
	}

	return rs
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"` //tag
	Name            string           `json:"name" gorm:"column:name;"`
	Addr            string           `json:"addr" gorm:"column:addr;"`
	UserId          int              `json:"-" gorm:"column:user_id"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
