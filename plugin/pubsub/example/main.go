package main

func main() {
	// var localPb pubsub.Pubsub = pblocal.NewPubSub("test1")
	// var topic = "OrderCreated"

	// sub1, close1 := localPb.Subscribe(context.Background(), topic)

	// sub2, _ := localPb.Subscribe(context.Background(), topic)

	// localPb.Publish(context.Background(), topic, pubsub.NewMessage(1))
	// localPb.Publish(context.Background(), topic, pubsub.NewMessage(2))

	// go func() {
	// 	for {
	// 		log.Println("Sub1: ", (<-sub1).Data())
	// 		time.Sleep(time.Millisecond * 400)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		log.Println("Sub2: ", (<-sub2).Data())
	// 		time.Sleep(time.Millisecond * 400)
	// 	}
	// }()

	// time.Sleep(time.Second * 3)
	// close1()

	// localPb.Publish(context.Background(), topic, pubsub.NewMessage(3))

	// time.Sleep(time.Second * 2)

}
