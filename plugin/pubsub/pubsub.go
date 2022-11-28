package pubsub

type Topic string
type Pubsub interface {
	Publish(done chan bool, exchangeName string, routingKey string, data *Message) error
	Subscribe(done chan error, exchangeName string, queueName string, routingKey string) (ch <-chan *Message, close func())
	//UnSubcribe(ctx context.Context, channel Channel) error
}
