package restaurantstorage

import (
	"context"
	restaurantmodel "food_delivery/module/restaurant/model"
)

func (store *sqlStore) DeleteRestaurant(
	ctx context.Context,
	cond map[string]interface{},
) error {

	if err := store.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
