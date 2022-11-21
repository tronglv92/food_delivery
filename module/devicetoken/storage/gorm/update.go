package devicetokenstorage

import (
	"context"
	"food_delivery/common"
	devicetokenmodel "food_delivery/module/devicetoken/model"
)

func (store *sqlStore) UpdateDeviceToken(
	ctx context.Context,
	cond map[string]interface{},
	data *devicetokenmodel.UserDeviceTokenUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
