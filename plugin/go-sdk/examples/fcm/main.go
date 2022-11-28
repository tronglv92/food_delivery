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
	deviceToken = "cYByoEGcR9SNpFAhGy4dan:APA91bH_HkTS6OwCxXk4o-QLBqZVTLuLHnvMtvS08s6XQzo_MOlhG2b_8TsXnDJ9LuJD8n5raMOjb2I4v-uki8p0bjJoOH2fq6d_K7ewH4Cvm3jd4Zt6nLNFKW8-hoS6ossV8ZYfGiRq"
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
