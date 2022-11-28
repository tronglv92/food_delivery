package restaurantrepo

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	"food_delivery/plugin/go-sdk/logger"
	"food_delivery/plugin/pubsub"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}
type DeviceTokenStore interface {
	GetDeviceTokens(ctx context.Context, userId int32) ([]common.SimpleDeviceToken, error)
}
type createRestaurantRepo struct {
	store       CreateRestaurantStore
	deviceToken DeviceTokenStore
	ps          pubsub.Pubsub
}

func NewCreateRestaurantRepo(store CreateRestaurantStore, deviceToken DeviceTokenStore, ps pubsub.Pubsub) *createRestaurantRepo {
	return &createRestaurantRepo{store: store, deviceToken: deviceToken, ps: ps}
}
func (repo *createRestaurantRepo) CreateRestaurant(context context.Context,
	data *restaurantmodel.RestaurantCreate) error {
	logger := logger.GetCurrent().GetLogger("restaurant.repository.create_restaurant")
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	if err := repo.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	deviceToken, err := repo.deviceToken.GetDeviceTokens(context, int32(data.UserId))
	if err != nil {
		logger.Errorln(err)
	}
	logger.Infof("deviceToken ", deviceToken)

	fcmTokens := make([]string, len(deviceToken))

	for _, item := range deviceToken {
		logger.Debugf("item.Token ", item)
		fcmTokens = append(fcmTokens, item.Token)
	}

	// newMessage := pubsub.NewMessage(map[string]interface{}{
	// 	"device_tokens": fcmTokens,
	// })
	// _ = repo.ps.Publish(context, common.TopicSendNotification, newMessage)
	return nil
}
