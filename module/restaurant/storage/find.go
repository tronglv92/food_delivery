package restaurantstorage

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindRestaurant(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
