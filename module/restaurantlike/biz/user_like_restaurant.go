package rstlikebiz

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	"food_delivery/plugin/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	CheckUserLike(ctx context.Context, userId, restaurantId int) (bool, error)
}

type IncLikeCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// incStore IncLikeCountResStore
	ps pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	// incStore IncLikeCountResStore,
	ps pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		// incStore: incStore,
		ps: ps,
	}
}
func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	liked, err := biz.store.CheckUserLike(ctx, data.UserId, data.RestaurantId)
	if err != nil && err != common.ErrRecordNotFound {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	if liked {
		return restaurantlikemodel.ErrUserAlreadyLikedRestaurant(nil)
	}

	err = biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// // Side effect
	// go func() {
	// 	defer common.AppRecover()
	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	// 			logger.GetCurrent().GetLogger("user.like.restaurant").Errorln(err)
	// 			return err
	// 		}

	// 		return nil
	// 	}, asyncjob.WithName("IncreaseLikeCount"))

	// 	if err := asyncjob.NewGroup(false, job).Run(ctx); err != nil {
	// 		logger.GetCurrent().GetLogger("user.like.restaurant").Errorln(err)
	// 	}
	// }()
	// newMessage := pubsub.NewMessage(data)
	newMessage := pubsub.NewMessage(map[string]interface{}{
		"user_id":       data.UserId,
		"restaurant_id": data.RestaurantId,
	})
	_ = biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, newMessage)

	return nil
}
