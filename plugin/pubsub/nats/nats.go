package appnats

import (
	"context"
	"encoding/json"
	"flag"
	"food_delivery/plugin/pubsub"
	"time"

	"food_delivery/plugin/go-sdk/logger"

	"github.com/nats-io/nats.go"
)

type appNats struct {
	name       string
	connection *nats.Conn
	logger     logger.Logger
	url        string
}

func NewNATS(name string) *appNats {
	return &appNats{
		name: name,
	}
}
func (an *appNats) GetPrefix() string {
	return an.name
}
func (an *appNats) Get() interface{} {
	return an
}
func (an *appNats) Name() string {
	return an.name
}
func (an *appNats) InitFlags() {
	flag.StringVar(&an.url, an.name+"-url", nats.DefaultURL, "URL of NATS service")
}
func (an *appNats) Configure() error {
	an.logger = logger.GetCurrent().GetLogger(an.name)

	an.logger.Infoln("Connecting to NATS service...")
	conn, err := nats.Connect(an.url, an.setupConnOptions([]nats.Option{})...)

	if err != nil {
		an.logger.Fatalln(err)
	}

	an.logger.Infoln("Connected to NATS service.")

	an.connection = conn

	return nil
}

func (an *appNats) Run() error {
	return an.Configure()
}

func (an *appNats) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
func (an *appNats) setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		an.logger.Infof("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		an.logger.Infof("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		an.logger.Infof("Exiting: %v", nc.LastError())
	}))
	return opts
}
func (an *appNats) Publish(ctx context.Context, channel string, data *pubsub.Message) error {
	msgData, err := json.Marshal(data.Data())

	if err != nil {
		an.logger.Errorln(err)
		return err
	}

	if err := an.connection.Publish(channel, msgData); err != nil {
		an.logger.Errorln(err)
		return err
	}

	return nil
}
func (an *appNats) Subscribe(ctx context.Context, channel string) (ch <-chan *pubsub.Message, close func()) {
	msgChan := make(chan *pubsub.Message)

	//go func() {}()
	sub, err := an.connection.Subscribe(channel, func(msg *nats.Msg) {
		msgData := make(map[string]interface{})

		_ = json.Unmarshal(msg.Data, &msgData)

		appMsg := pubsub.NewMessage(msgData)
		appMsg.SetChannel(channel)

		msgChan <- appMsg

	})

	if err != nil {
		an.logger.Errorln(err)
	}

	return msgChan, func() {
		_ = sub.Unsubscribe()
	}
}
