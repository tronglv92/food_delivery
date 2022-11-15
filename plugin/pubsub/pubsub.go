package pubsub

import "context"

type Topic string
type Pubsub interface {
	Publish(ctx context.Context, channel string, data *Message) error
	Subscribe(ctx context.Context, channel string) (ch <-chan *Message, close func())
	//UnSubcribe(ctx context.Context, channel Channel) error
}
