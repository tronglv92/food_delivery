package devicetokenstorage

import (
	"context"
	devicetokenmodel "food_delivery/module/devicetoken/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDeviceToken(ctx context.Context, userId int, deviceId string) *devicetokenmodel.UserDeviceToken {
	var data devicetokenmodel.UserDeviceToken

	if err := s.db.
		Where("user_id = ? and device_id = ?", userId, deviceId).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return nil
	}

	return &data
}
