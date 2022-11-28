package grpcstore

import (
	"context"
	"food_delivery/common"
	user "food_delivery/proto"
)

type UserStorage interface {
	// user/storage/gorm/get: GetUsers
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}
type userGRPCBusiness struct {
	dbStore UserStorage
}

func NewUserGRPCBusiness(dbStore UserStorage) *userGRPCBusiness {
	return &userGRPCBusiness{dbStore: dbStore}
}

func (s *userGRPCBusiness) GetUserByIds(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	userIds := make([]int, len(request.GetUserIds()))
	for i := range userIds {
		userIds[i] = int(request.GetUserIds()[i])
	}

	rs, err := s.dbStore.GetUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, len(rs))

	for i, item := range rs {
		item.Mask(common.DbTypeUser)

		users[i] = &user.User{
			Id:        item.FakeId.String(),
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Role:      item.Role,
		}
	}
	return &user.UserResponse{Users: users}, nil
}
