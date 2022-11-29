package userbiz

import (
	"context"

	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
)

type RenewTokenRepo interface {
	RenewAccessToken(ctx context.Context, data *usermodel.RenewAccessTokenRequest) (*usermodel.Account, error)
}

type renewTokenBiz struct {
	repo RenewTokenRepo
}

func NewRenewTokenBiz(repo RenewTokenRepo) *renewTokenBiz {
	return &renewTokenBiz{repo: repo}
}
func (biz *renewTokenBiz) RenewAccessToken(ctx context.Context, data *usermodel.RenewAccessTokenRequest) (*usermodel.Account, error) {
	account, err := biz.repo.RenewAccessToken(ctx, data)
	if err != nil {
		return nil, common.NewCustomError(
			err,
			"Cannot Renew Token",
			"ErrCanNotRenewToken",
		)
	}

	// newMessage := pubsub.NewMessage(map[string]interface{}{
	// 	"user_id":       data.UserId,
	// 	"restaurant_id": data.RestaurantId,
	// })
	// _ = biz.ps.Publish(context, common.TopicUserLikeRestaurant, newMessage)
	return account, nil
}
