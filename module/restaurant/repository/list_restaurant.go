package restaurantrepo

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
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
	store  ListRestaurantRepo
	uStore UserStore
	// likeStore LikeRestaurantStore
}

type UserStore interface {
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

func NewListRestaurantRepo(store ListRestaurantRepo, uStore UserStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, uStore: uStore}
}
func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	// logger := logger.GetCurrent().GetLogger("restaurant.repo.list_restaurant")
	// result, err := repo.store.ListDataWithCondition(ctx, filter, paging, "User")
	result, err := repo.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	// logger.Debugf("vao trong nay", result)
	// fmt.Println("restaurant.repo.list_restaurant: ", result)
	userIds := make([]int, len(result))

	for i := range userIds {
		userIds[i] = result[i].UserId
	}

	users, err := repo.uStore.GetUsers(ctx, userIds)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)
	// if err != nil {
	// 	log.Println(err)
	// 	return result, nil
	// }

	// for i, item := range result {
	// 	result[i].LikedCount = likeMap[item.Id]
	// }

	// O(N^2)
	//for i := range result {
	//	for j := range users {
	//		if result[i].OwnerId == users[j].Id {
	//			result[i].Owner = &users[j]
	//			break
	//		}
	//	}
	//}

	// O(N)
	mapUser := make(map[int]*common.SimpleUser)

	for j, u := range users {
		mapUser[u.Id] = &users[j]
	}

	for i, item := range result {
		result[i].User = mapUser[item.UserId]
	}

	return result, nil
}
