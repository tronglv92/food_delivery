package appgrpc

import (
	"context"
	"food_delivery/common"
	"food_delivery/plugin/go-sdk/logger"
	user "food_delivery/proto"
)

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
