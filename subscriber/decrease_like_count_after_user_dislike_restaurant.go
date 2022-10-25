package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"food_delivery/pubsub"
)

// func DecreaseLikeCountAfterUserDisLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubSub().Subcribe(ctx, common.TopicUserDislikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)

// 			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }
func DecreaseLikeCountAfterUserDisLikeRestaurant(appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Decrease like count after user dislikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}
