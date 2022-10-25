package restaurantrepo

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	"log"
)

type ListRestaurantRepo interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}
type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}
type listRestaurantRepo struct {
	store     ListRestaurantRepo
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantRepo, likeStore LikeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, likeStore: likeStore}
}
func (biz *listRestaurantRepo) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range ids {
		ids[i] = result[i].Id
	}

	likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)
	if err != nil {
		log.Println(err)
		return result, nil
	}

	for i, item := range result {
		result[i].LikedCount = likeMap[item.Id]
	}

	return result, nil
}
