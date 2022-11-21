package devicetokenbiz

import (
	"context"
	"fmt"
	"food_delivery/common"
	devicetokenmodel "food_delivery/module/devicetoken/model"
)

type PutDeviceTokenStore interface {
	Create(context context.Context, data *devicetokenmodel.UserDeviceTokenUpdate) error
	FindDeviceToken(ctx context.Context, userId int, deviceId string) *devicetokenmodel.UserDeviceToken
	UpdateDeviceToken(
		ctx context.Context,
		cond map[string]interface{},
		data *devicetokenmodel.UserDeviceTokenUpdate,
	) error
}
type putDeviceTokenBiz struct {
	store PutDeviceTokenStore
}

func NewPutDeviceTokenBiz(store PutDeviceTokenStore) *putDeviceTokenBiz {
	return &putDeviceTokenBiz{store: store}
}
func (biz *putDeviceTokenBiz) CreateDeviceToken(context context.Context,
	data *devicetokenmodel.UserDeviceTokenUpdate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	fcm := biz.store.FindDeviceToken(context, data.UserId, data.DeviceId)

	fmt.Println("fcm ", fcm)
	if fcm != nil {
		if err := biz.store.UpdateDeviceToken(context, map[string]interface{}{"id": fcm.Id}, data); err != nil {
			return common.ErrCannotCreateEntity(devicetokenmodel.EntityName, err)
		}
		return nil
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(devicetokenmodel.EntityName, err)
	}
	return nil
}
