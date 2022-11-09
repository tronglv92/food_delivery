package appgrpc

import (
	"context"
	"flag"
	"food_delivery/common"
	user "food_delivery/proto"

	"food_delivery/plugin/go-sdk/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	client      user.UserServiceClient
}

func NewUserClient(prefix string) *grpcClient {
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

	uc.client = user.NewUserServiceClient(cc)

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

func (uc *grpcClient) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	logger.GetCurrent().GetLogger(uc.prefix).Infoln("GetUsers grpc store running")

	userIds := make([]int32, len(ids))

	for i := range userIds {
		userIds[i] = int32(ids[i])
	}

	rs, err := uc.client.GetUserByIds(ctx, &user.UserRequest{UserIds: userIds})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(rs.Users))

	for i, item := range rs.Users {
		uid, _ := common.FromBase58(item.Id)

		users[i] = common.SimpleUser{
			SQLModel: common.SQLModel{
				Id:     int(uid.GetLocalID()),
				FakeId: &uid,
			},
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Role:      item.Role,
		}
	}

	return users, nil
}
