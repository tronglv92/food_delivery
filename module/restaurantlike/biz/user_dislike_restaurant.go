package rstlikebiz

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"food_delivery/pubsub"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

// type DecLikeCountResStore interface {
// 	DecreaseLikeCount(ctx context.Context, id int) error
// }

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
	// decStore DecLikeCountResStore
	ps pubsub.Pubsub
}

func NewUserDislikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	// decStore DecLikeCountResStore,
	ps pubsub.Pubsub,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
		// decStore: decStore,
		ps: ps,
	}
}
func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	// Get Like count of restaurant => <0 thi khong decrease

	data := &restaurantlikemodel.Like{RestaurantId: restaurantId}
	// Send message
	if err := biz.ps.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}
	// // Side effect
	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	return nil

}
