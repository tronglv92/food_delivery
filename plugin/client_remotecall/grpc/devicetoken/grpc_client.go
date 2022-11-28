package appgrpc

import (
	"flag"
	user "food_delivery/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	client      user.DeviceTokenServiceClient
}

func NewDeviceTokenClient(prefix string) *grpcClient {
	return &grpcClient{
		prefix: prefix,
	}
}

func (uc *grpcClient) GetPrefix() string {
	return uc.prefix
}

func (uc *grpcClient) Get() interface{} {
	return uc
}

func (uc *grpcClient) Name() string {
	return uc.prefix
}

func (uc *grpcClient) InitFlags() {
	flag.StringVar(&uc.url, uc.GetPrefix()+"-url", "localhost:50051", "URL connect to grpc server")
}

func (uc *grpcClient) Configure() error {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(uc.url, opts)

	if err != nil {
		return err
	}

	uc.client = user.NewDeviceTokenServiceClient(cc)

	return nil
}

func (uc *grpcClient) Run() error {
	return uc.Configure()
}

func (uc *grpcClient) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		c <- true
	}()
	return c
}
