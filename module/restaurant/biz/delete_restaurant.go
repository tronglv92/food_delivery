package restaurantbiz

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	DeleteRestaurant(
		ctx context.Context,
		cond map[string]interface{},
	) error
	FindRestaurant(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
	// requester common.Requester
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}
func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindRestaurant(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	// if oldData.UserId != biz.requester.GetUserId() {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.DeleteRestaurant(context, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil)
	}
	return nil
}
