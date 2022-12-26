package asynqclient

import (
	"flag"
	"food_delivery/plugin/go-sdk/logger"
	"log"

	"github.com/hibiken/asynq"
)

type AsynqClient struct {
	prefix    string
	logger    logger.Logger
	client    *asynq.Client
	redisAddr string
}

func NewAsynqClient(prefix string) *AsynqClient {

	return &AsynqClient{

		prefix: prefix,
	}
}

func (s *AsynqClient) Get() interface{} {
	return s
}

func (s *AsynqClient) Name() string {
	return s.prefix
}
func (s *AsynqClient) InitFlags() {
	flag.StringVar(&s.redisAddr, s.prefix+"-redis-addr", "127.0.0.1:6379", "Address Redis of asynq Client")
	flag.Parse()
}
func (s *AsynqClient) GetPrefix() string {
	return s.prefix
}
func (s *AsynqClient) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.Name())

	return nil
}

func (s *AsynqClient) Run() error {
	return s.Configure()
}
func (s *AsynqClient) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (s *AsynqClient) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	log.Println("vao trong nay Enqueue")
	s.client = asynq.NewClient(asynq.RedisClientOpt{Addr: s.redisAddr})
	defer s.client.Close()

	info, err := s.client.Enqueue(task, opts...)
	if err != nil {
		return nil, err
	}
	return info, nil
}
