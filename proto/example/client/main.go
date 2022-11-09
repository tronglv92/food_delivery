package main

import (
	"context"
	user "food_delivery/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()
	client := user.NewUserServiceClient(cc)

	for i := 1; i <= 5; i++ {
		res, err := client.GetUserByIds(context.Background(), &user.UserRequest{UserIds: []int32{1}})

		if err != nil {
			log.Println(err)
		} else {
			log.Println(res.Users)
		}

		time.Sleep(time.Second * 6)
	}
}
