package cmd

import (
	"fmt"
	"food_delivery/common"
	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/pubsub"
	rabbitmq "food_delivery/plugin/pubsub/rabbitmq"
	"food_delivery/plugin/storage/sdkgorm"
	"log"

	"github.com/spf13/cobra"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

var startSubUserLikedRestaurantCmd = &cobra.Command{
	Use:   "sub-user-liked-restaurant",
	Short: "Start a subscriber when user liked restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
			goservice.WithInitRunnable(rabbitmq.NewRabbitMQ(common.PluginRabbitMQ)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		ps := service.MustGet(common.PluginRabbitMQ).(pubsub.Pubsub)

		// ctx := context.Background()

		done := make(chan error)
		ch, _ := ps.Subscribe(done, "test-exchange", "test-queue", "test-key")

		for msg := range ch {
			// db := service.MustGet(common.DBMain).(*gorm.DB)

			// if likeData, ok := msg.Data().(HasRestaurantId); ok {
			// 	job := asyncjob.NewJob(func(ctx context.Context) error {
			// 		return restaurantstorage.NewSQLStore(db).IncreaseLikeCount(ctx, likeData.GetRestaurantId())
			// 	})

			// 	if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
			// 		log.Println(err)
			// 	}
			// }

			if restaurantId, ok := msg.Data["restaurant_id"]; ok {
				fmt.Printf("restaurantId ", restaurantId)
				// job := asyncjob.NewJob(func(ctx context.Context) error {
				// 	return restaurantstorage.NewSQLStore(db).IncreaseLikeCount(ctx, int(restaurantId.(float64)))
				// })

				// if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
				// 	log.Println(err)
				// }
			}
		}
	},
}
