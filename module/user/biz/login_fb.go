package userbiz

import (
	"context"
	"food_delivery/common"

	usermodel "food_delivery/module/user/model"
)

type LoginFacebookRepo interface {
	LoginFacebook(ctx context.Context, accessToken string) (*usermodel.Account, error)
}

type loginFacebookBiz struct {
	repo LoginFacebookRepo
}

func NewLoginFacebookBiz(repo LoginFacebookRepo) *loginFacebookBiz {
	return &loginFacebookBiz{repo: repo}
}
func (biz *loginFacebookBiz) LoginFacebook(ctx context.Context, accessToken string) (*usermodel.Account, error) {
	user, err := biz.repo.LoginFacebook(ctx, accessToken)
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
