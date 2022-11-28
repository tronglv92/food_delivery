package cmd

import (
	"context"
	"food_delivery/plugin/fcm"

	"github.com/spf13/cobra"
)

var (
	deviceToken = "c_GPTomvTpWGXiU-Y83w4y:APA91bHDrHpIQai5zgeWVghYEtfKEgWa2PlXPcPkXRdZhMy9D_RVnXSP2A5qN-T-2oS3IUb4s664yzZ-ukBbrJwpdf30q38QEzIAixZbfx9b53SFwwptv4SJpqM4Z_igBL-8HZ1I3YVu"
)
var startSubSendNotificationCmd = &cobra.Command{
	Use:   "sub-send-notification",
	Short: "Topic Send notification",
	Run: func(cmd *cobra.Command, args []string) {
		// service := goservice.New(
		// 	goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		// 	goservice.WithInitRunnable(appnats.NewNATS(common.PluginNATS)),
		// 	goservice.WithInitRunnable(fcm.New(common.PluginFCM)),
		// )

		// if err := service.Init(); err != nil {
		// 	log.Fatalln(err)
		// }

		// ps := service.MustGet(common.PluginNATS).(pubsub.Pubsub)

		// ctx := context.Background()

		// ch, _ := ps.Subscribe(ctx, common.TopicSendNotification)
		// for msg := range ch {

		// 	if results, ok := msg.Data()["device_tokens"]; ok {

		// 		rsInterface := results.([]interface{})
		// 		fcmTokens := make([]string, len(rsInterface))
		// 		for i, v := range rsInterface {
		// 			fcmTokens[i] = v.(string)
		// 			fmt.Printf("deviceTokens ", fcmTokens[i])
		// 		}

		// 		job := asyncjob.NewJob(func(ctx context.Context) error {

		// 			client := service.MustGet(common.PluginFCM).(fcm.FirebaseCloudMessaging)
		// 			client.ShowPrintResult(true)

		// 			collapseKey := "messages is my name"
		// 			notification := fcm.NewNotification("title",
		// 				fcm.WithIcon("ic_notification"),
		// 				fcm.WithColor("#18d821"),
		// 				fcm.WithSound("default"),
		// 				fcm.WithCollapseKey(collapseKey),
		// 				fcm.WithTag(collapseKey),
		// 			)
		// 			sendToDeviceList(client, notification, fcmTokens)
		// 			return nil
		// 		})

		// 		if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
		// 			log.Println(err)
		// 		}
		// 	}
		// }
	},
}

// func sendToTopic(client fcm.FirebaseCloudMessaging, notification *fcm.Notification) {
// 	res, err := client.SendToTopic(context.Background(), "/topics/chat", notification)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(res)
// }
// func sendToDevice(client fcm.FirebaseCloudMessaging, notification *fcm.Notification) {
// 	res, err := client.SendToDevice(context.Background(), deviceToken, notification)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(res)
// }
func sendToDeviceList(client fcm.FirebaseCloudMessaging, notification *fcm.Notification, deviceTokens []string) (*fcm.Response, error) {
	res, err := client.SendToDevices(context.Background(), deviceTokens, notification)
	if err != nil {
		return nil, err
	}
	return res, nil
}
