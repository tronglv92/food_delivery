package pblocal

import (
	"context"
	"food_delivery/common"
	"food_delivery/pubsub"
	"log"
	"sync"
)

// A pb run locally (in-mem)
// It has a queue (buffer channel) at it's core and many group of subscribers
// Because we want to send a message with a specific topic for many subscribers in a group cai gi do

type localPubSub struct {
	name         string
	messageQueue chan *pubsub.Message

	mapChannel map[string][]chan *pubsub.Message
	locker     *sync.RWMutex
}

func NewPubSub(name string) *localPubSub {
	pb := &localPubSub{
		name:         name,
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[string][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb

}

func (ps *localPubSub) Publish(ctx context.Context, topic string, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published", data.String(), "with data", data.Data())
	}()
	return nil
}
func (ps *localPubSub) Subscribe(ctx context.Context, topic string) (ch <-chan *pubsub.Message, unsubscribe func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					// [1,2,3,4,5] //  i = 3
					// [1,2,3] (arr[:i])
					// [5] (arr[i+1:])
					// [1,2,3,5]
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()

					close(c)
					break

				}
			}
		}
	}
}
func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue", mess)

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
func (ps *localPubSub) GetPrefix() string {
	return ps.name
}

func (ps *localPubSub) Get() interface{} {
	return ps
}

func (ps *localPubSub) Name() string {
	return ps.name
}

func (ps *localPubSub) InitFlags() {
}

func (ps *localPubSub) Configure() error {
	return nil
}

func (ps *localPubSub) Run() error {
	return nil
}

func (ps *localPubSub) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
