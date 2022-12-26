package cmd

import (
	"context"
	"fmt"
	"food_delivery/common"
	tasksModel "food_delivery/module/sendtask/model"
	asynqserver "food_delivery/plugin/asynqserver"
	goservice "food_delivery/plugin/go-sdk"
	rabbitmq "food_delivery/plugin/pubsub/rabbitmq"
	"food_delivery/plugin/storage/sdkgorm"
	"food_delivery/subscribes"
	"log"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

var startSubEmailCmd = &cobra.Command{
	Use:   "sub-email",
	Short: "Start a job email",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(

			goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
			goservice.WithInitRunnable(asynqserver.NewAsynqServer(common.PluginAsynqServer)),
			goservice.WithInitRunnable(rabbitmq.NewRabbitMQ(common.PluginRabbitMQ)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		srv := service.MustGet(common.PluginAsynqServer).(*asynq.Server)
		if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
			log.Fatal(err)
		}
	},
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case tasksModel.TypeEmailDelivery:
		subscribes.HandleEmailDeliveryTask(ctx, t)

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}
