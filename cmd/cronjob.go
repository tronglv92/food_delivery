package cmd

import (
	cronjob2 "food_delivery/plugin/cronjob"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/spf13/cobra"
)

var cronjob = &cobra.Command{
	Use:   "cronjob",
	Short: "Run my cron job",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(cronjob2.NewMyCronJob()),
		)

		service.Init()

		service.Start()
	},
}
