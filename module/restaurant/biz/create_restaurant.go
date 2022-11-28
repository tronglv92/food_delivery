package restaurantbiz

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	
)

type CreateRestaurantRepo interface {
	CreateRestaurant(context context.Context,
		data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	repo CreateRestaurantRepo
}

func NewCreateRestaurantBiz(repo CreateRestaurantRepo) *createRestaurantBiz {
	return &createRestaurantBiz{repo: repo}
}
func (biz *createRestaurantBiz) CreateRestaurant(context context.Context,
	data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	if err := biz.repo.CreateRestaurant(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	// newMessage := pubsub.NewMessage(map[string]interface{}{
	// 	"user_id":       data.UserId,
	// 	"restaurant_id": data.RestaurantId,
	// })
	// _ = biz.ps.Publish(context, common.TopicUserLikeRestaurant, newMessage)
	return nil
}
