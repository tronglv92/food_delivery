package devicetokengrpcstorage

import (
	"context"
	"food_delivery/common"
	devicetokenmodel "food_delivery/module/devicetoken/model"
	user "food_delivery/proto"
)

type DeviceTokenStore interface {
	GetDeviceTokens(ctx context.Context, userId int) ([]devicetokenmodel.UserDeviceToken, error)
}
type deviceTokenGRPCBusiness struct {
	dbStore DeviceTokenStore
}

func NewDeviceTokenGRPCBusiness(dbStore DeviceTokenStore) *deviceTokenGRPCBusiness {
	return &deviceTokenGRPCBusiness{dbStore: dbStore}
}
func (s *deviceTokenGRPCBusiness) GetDeviceTokenId(ctx context.Context, request *user.DeviceTokenRequest) (*user.DeviceTokenResponse, error) {

	rs, err := s.dbStore.GetDeviceTokens(ctx, int(request.GetUserId()))
	if err != nil {
		return nil, err
	}

	deviceTokens := make([]*user.DeviceToken, len(rs))

	for i, item := range rs {
		item.Mask(common.DbTypeUser)

		deviceTokens[i] = &user.DeviceToken{
			Id:           item.FakeId.String(),
			UserId:       int32(item.UserId),
			IsProduction: int32(item.IsProduction),
			Os:           item.OS,
			DeviceId:     item.DeviceId,
			Token:        item.Token,
		}
	}
	return &user.DeviceTokenResponse{DeviceTokens: deviceTokens}, nil
}
