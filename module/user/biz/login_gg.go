package userbiz

import (
	"context"
	"food_delivery/common"

	usermodel "food_delivery/module/user/model"
)

type LoginGoogleRepo interface {
	LoginGoogle(ctx context.Context, accessToken string) (*usermodel.Account, error)
}

type loginGoogleBiz struct {
	repo LoginGoogleRepo
}

func NewLoginGoogleBiz(repo LoginGoogleRepo) *loginGoogleBiz {
	return &loginGoogleBiz{repo: repo}
}
func (biz *loginGoogleBiz) LoginGoogle(ctx context.Context, accessToken string) (*usermodel.Account, error) {
	user, err := biz.repo.LoginGoogle(ctx, accessToken)
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
