package asynqserver

import (
	"flag"
	"food_delivery/plugin/go-sdk/logger"

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

	s.Configure()
	// if err := s.server.Run(asynq.HandlerFunc(handler)); err != nil {
	// 	log.Fatal(err)
	// }
	return nil

}
func (s *AsynqServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
