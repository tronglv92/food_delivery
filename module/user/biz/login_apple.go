package userbiz

import (
	"context"
	"food_delivery/common"

	usermodel "food_delivery/module/user/model"
)

type LoginAppleRepo interface {
	LoginApple(ctx context.Context, accessToken string) (*usermodel.Account, error)
}

type loginAppleBiz struct {
	repo LoginAppleRepo
}

func NewLoginAppleBiz(repo LoginAppleRepo) *loginAppleBiz {
	return &loginAppleBiz{repo: repo}
}
func (biz *loginAppleBiz) LoginApple(ctx context.Context, accessToken string) (*usermodel.Account, error) {
	user, err := biz.repo.LoginApple(ctx, accessToken)
	if err != nil {
		return nil, common.ErrLoginNotValid(err)
	}

	// newMessage := pubsub.NewMessage(map[string]interface{}{
	// 	"user_id":       data.UserId,
	// 	"restaurant_id": data.RestaurantId,
	// })
	// _ = biz.ps.Publish(context, common.TopicUserLikeRestaurant, newMessage)
	return user, nil
}
