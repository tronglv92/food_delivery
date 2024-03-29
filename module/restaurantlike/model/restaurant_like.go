package restaurantlikemodel

import (
	"fmt"
	"food_delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string { return "restaurant_likes" }
func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}
func (l *Like) GetUserId() int {
	return l.UserId
}
func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}
func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this restaurant"),
		fmt.Sprintf("ErrCannotUnLikeRestaurant"),
	)
}
func ErrUserAlreadyLikedRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("you've already liked this restaurant"),
		fmt.Sprintf("ErrUserAlreadyLikedRestaurant"),
	)
}
func ErrCannotDidNotlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You have not liked this restaurant"),
		fmt.Sprintf("ErrCannotDidNotlikeRestaurant"),
	)
}
