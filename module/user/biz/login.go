package userbiz

import (
	"context"
	"food_delivery/common"

	usermodel "food_delivery/module/user/model"
)

type LoginRepo interface {
	Login(ctx context.Context, userAgent string, clientIp string, data *usermodel.UserLogin) (*usermodel.Account, error)
}

type loginBiz struct {
	repo LoginRepo
}

func NewLoginBiz(repo LoginRepo) *loginBiz {
	return &loginBiz{repo: repo}
}
func (biz *loginBiz) Login(ctx context.Context, userAgent string, clientIp string, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.repo.Login(ctx, userAgent, clientIp, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	// newMessage := pubsub.NewMessage(map[string]interface{}{
	// 	"user_id":       data.UserId,
	// 	"restaurant_id": data.RestaurantId,
	// })
	// _ = biz.ps.Publish(context, common.TopicUserLikeRestaurant, newMessage)
	return user, nil
}
