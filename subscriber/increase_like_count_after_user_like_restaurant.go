package subscriber

import (
	"context"
	"fmt"
	"food_delivery/component/appctx"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"food_delivery/pubsub"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}

}
func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Push notificaiton when user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user likes restaurant id:", likeData.GetRestaurantId())
			return nil
		},
	}

}
func EmitRealtimeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {

	return consumerJob{
		Title: "Realtime emit after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			fmt.Println("likeData ", likeData)
			fmt.Println("GetRealtimeEngine() ", appCtx.GetRealtimeEngine())
			appCtx.GetRealtimeEngine().EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)
			return nil
		},
	}

}
