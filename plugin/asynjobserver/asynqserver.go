package asynjobserver

import (
	"context"
	"flag"
	"fmt"
	tasksModel "food_delivery/module/sendtask/model"
	"food_delivery/plugin/go-sdk/logger"
	"log"

	"food_delivery/subscribes"

	"github.com/hibiken/asynq"
)

type AsynqServer struct {
	prefix string
	logger logger.Logger
	server *asynq.Server

	redisAddr string
}

func NewAsynqServer(prefix string) *AsynqServer {

	return &AsynqServer{

		prefix: prefix,
	}
}
func (s *AsynqServer) Get() interface{} {
	return s.server
}

func (s *AsynqServer) Name() string {
	return s.prefix
}
func (s *AsynqServer) InitFlags() {
	flag.StringVar(&s.redisAddr, s.prefix+"-redis-addr", "127.0.0.1:6379", "Address Redis of asynq Client")
	flag.Parse()
}
func (s *AsynqServer) GetPrefix() string {
	return s.prefix
}
func (s *AsynqServer) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.Name())
	s.server = asynq.NewServer(
		asynq.RedisClientOpt{Addr: s.redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	return nil
}

func (s *AsynqServer) Run() error {

	go func() {
		s.Configure()
		if err := s.server.Run(asynq.HandlerFunc(handler)); err != nil {
			log.Fatal(err)
		}
	}()

	return nil

}
func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case tasksModel.TypeEmailDelivery:
		subscribes.HandleEmailDeliveryTask(ctx, t)

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}

func (s *AsynqServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
