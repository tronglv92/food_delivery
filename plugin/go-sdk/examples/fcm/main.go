package main

import (
	"context"
	// "context"
	"fmt"
	"food_delivery/plugin/fcm"
	goservice "food_delivery/plugin/go-sdk"
	"log"
)

var (
	deviceToken = "c_GPTomvTpWGXiU-Y83w4y:APA91bHDrHpIQai5zgeWVghYEtfKEgWa2PlXPcPkXRdZhMy9D_RVnXSP2A5qN-T-2oS3IUb4s664yzZ-ukBbrJwpdf30q38QEzIAixZbfx9b53SFwwptv4SJpqM4Z_igBL-8HZ1I3YVu"
)

func main() {
	service := goservice.New(
		goservice.WithName("demo"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(fcm.New("fcm")),
	)
	_ = service.Init()

	client := service.MustGet("fcm").(fcm.FirebaseCloudMessaging)
	client.ShowPrintResult(true)

	collapseKey := "messages is my name"

	notification := fcm.NewNotification("title",
		fcm.WithIcon("ic_notification"),
		fcm.WithColor("#18d821"),
		fcm.WithSound("default"),
		fcm.WithCollapseKey(collapseKey),
		fcm.WithTag(collapseKey),
	)

	sendToDeviceList(client, notification)
	// sendToDevice(client, notification)
	// sendToTopic(client, notification)
}

func sendToDevice(client fcm.FirebaseCloudMessaging, notification *fcm.Notification) {
	res, err := client.SendToDevice(context.Background(), deviceToken, notification)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func sendToDeviceList(client fcm.FirebaseCloudMessaging, notification *fcm.Notification) {
	res, err := client.SendToDevices(context.Background(), []string{"xxx", deviceToken, "123"}, notification)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func sendToTopic(client fcm.FirebaseCloudMessaging, notification *fcm.Notification) {
	res, err := client.SendToTopic(context.Background(), "/topics/chat", notification)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
