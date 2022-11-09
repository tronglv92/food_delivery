package cmd

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/asyncjob"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"food_delivery/plugin/sdkgorm"
	"food_delivery/pubsub"
	appnats "food_delivery/pubsub/nats"
	"log"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var startSubUserDislikedRestaurantCmd = &cobra.Command{
	Use:   "sub-user-disliked-restaurant",
	Short: "Start a subscriber when user disliked restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
			goservice.WithInitRunnable(appnats.NewNATS(common.PluginNATS)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		ps := service.MustGet(common.PluginNATS).(pubsub.Pubsub)

		ctx := context.Background()

		ch, _ := ps.Subscribe(ctx, common.TopicUserDislikeRestaurant)

		for msg := range ch {
			db := service.MustGet(common.DBMain).(*gorm.DB)

			if restaurantId, ok := msg.Data()["restaurant_id"]; ok {
				job := asyncjob.NewJob(func(ctx context.Context) error {
					return restaurantstorage.NewSQLStore(db).DecreaseLikeCount(ctx, int(restaurantId.(float64)))
				})

				if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
					log.Println(err)
				}
			}
		}

		// for msg := range ch {
		// 	db := service.MustGet(common.DBMain).(*gorm.DB)

		// 	if dislikedData, ok := msg.Data().(HasRestaurantId); ok {
		// 		job := asyncjob.NewJob(func(ctx context.Context) error {
		// 			return restaurantstorage.NewSQLStore(db).DecreaseLikeCount(ctx, dislikedData.GetRestaurantId())
		// 		})

		// 		if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
		// 			log.Println(err)
		// 		}
		// 	}
		// }
	},
}
