package rstlikebiz

import (
	"context"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"food_delivery/plugin/pubsub"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
	FindUserLike(ctx context.Context, userId, restaurantId int) (*restaurantlikemodel.Like, error)
}

type DecLikeCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

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
	oldData, err := biz.store.FindUserLike(ctx, userId, restaurantId)

	if oldData == nil {
		return restaurantlikemodel.ErrCannotDidNotlikeRestaurant(err)
	}

	err = biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	// // Side effect
	// go func() {
	// 	defer common.AppRecover()
	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		if err := biz.decStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
	// 			logger.GetCurrent().GetLogger("user.dislike.restaurant").Errorln(err)
	// 			return err
	// 		}

	// 		return nil
	// 	}, asyncjob.WithName("DecreaseLikeCount"))

	// 	if err := asyncjob.NewGroup(false, job).Run(ctx); err != nil {
	// 		logger.GetCurrent().GetLogger("user.dislike.restaurant").Errorln(err)
	// 	}
	// }()

	// PUBSUB
	// newMessage := pubsub.NewMessage(map[string]interface{}{
	// 	"user_id":       userId,
	// 	"restaurant_id": restaurantId,
	// })
	// biz.ps.Publish(ctx, common.TopicUserDislikeRestaurant, newMessage)

	return nil

}
