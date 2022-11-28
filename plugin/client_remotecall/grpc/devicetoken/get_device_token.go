package appgrpc

import (
	"context"
	"food_delivery/common"
	"food_delivery/plugin/go-sdk/logger"
	user "food_delivery/proto"
)

type SimpleDeviceToken struct {
	common.SQLModel
	Token        string `json:"token" gorm:"column:token;"`
	UserId       int    `json:"user_id" gorm:"column:user_id;"`
	DeviceId     string `json:"role" gorm:"column:role;"`
	IsProduction int    `json:"is_production" gorm:"column:is_production;"`

	//Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (uc *grpcClient) GetDeviceTokens(ctx context.Context, userId int32) ([]common.SimpleDeviceToken, error) {
	logger := logger.GetCurrent().GetLogger(uc.prefix)
	logger.Infoln("GetUsers grpc store running")

	rs, err := uc.client.GetDeviceTokenId(ctx, &user.DeviceTokenRequest{UserId: userId})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	deviceTokens := make([]common.SimpleDeviceToken, len(rs.DeviceTokens))

	for i, item := range rs.DeviceTokens {
		uid, _ := common.FromBase58(item.Id)

		logger.Debugf("item ", item)
		deviceTokens[i] = common.SimpleDeviceToken{
			SQLModel: common.SQLModel{
				Id:     int(uid.GetLocalID()),
				FakeId: &uid,
			},
			Token:        item.Token,
			UserId:       int(item.UserId),
			DeviceId:     item.DeviceId,
			IsProduction: int(item.IsProduction),
		}
	}

	return deviceTokens, nil
}
