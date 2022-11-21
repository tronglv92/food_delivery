package devicetokenstorage

import (
	"context"
	"food_delivery/common"
	devicetokenmodel "food_delivery/module/devicetoken/model"
)

func (s *sqlStore) Create(context context.Context, data *devicetokenmodel.UserDeviceTokenUpdate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
