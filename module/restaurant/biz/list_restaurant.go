package biz

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"

	"go.opencensus.io/trace"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}
func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	ctx1, span := trace.StartSpan(context, "biz.list_restaurant")
	result, err := biz.repo.ListRestaurant(ctx1, filter, paging)

	span.End()
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
