package devicetokenstorage

import (
	"context"
	devicetokenmodel "food_delivery/module/devicetoken/model"
)

func (s *sqlStore) GetDeviceTokens(ctx context.Context, userId int) ([]devicetokenmodel.UserDeviceToken, error) {
	var results []devicetokenmodel.UserDeviceToken
	var empty []devicetokenmodel.UserDeviceToken
	if err := s.db.Table(devicetokenmodel.UserDeviceToken{}.TableName()).
		Where("user_id = ?", userId).
		Find(&results).Error; err != nil {

		return empty, err
	}

	return results, nil
}
