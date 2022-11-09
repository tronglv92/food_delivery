package restaurantbiz

import (
	"context"
	"errors"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	"testing"
)

type smokeCreateStore struct{}

func (smokeCreateStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "test" {
		return common.ErrDB(errors.New("something went wrong in DB"))
	}
	data.Id = 200
	return nil
}
func TestNewCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurantBiz(smokeCreateStore{})
	dataTest := restaurantmodel.RestaurantCreate{Name: ""}
	err := biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil || err.Error() != "invalid request" {
		t.Errorf("Failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "test"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil {
		t.Errorf("Failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "test2"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err != nil {
		t.Errorf("Failed")
	}
}
