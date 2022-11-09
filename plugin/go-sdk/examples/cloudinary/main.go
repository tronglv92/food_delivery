package main

import (
	"context"
	"fmt"
	goservice "food_delivery/plugin/go-sdk"
	"food_delivery/plugin/go-sdk/plugin/cloudinary"
	"food_delivery/plugin/go-sdk/sdkcm"
	"log"
)

func main() {
	service := goservice.New(
		goservice.WithName("cloudinary"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(cloudinary.New("cloudinary")),
	)

	_ = service.Init()

	videoFile := "videotest.mov" // put this file on project root to test

	cloudinary := service.MustGet("cloudinary").(cloudinary.Cloudinary)

	result, err := cloudinary.VideoUpload(context.Background(), videoFile, "video_preset", "test", "mp4")

	if err != nil {
		log.Fatalf("err: %+v", err.(sdkcm.AppError).Log)
	}

	fmt.Printf("%+v", result)
}
