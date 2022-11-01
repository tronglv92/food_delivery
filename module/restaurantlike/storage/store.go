package restaurantlikestorage

import (
	"context"
	"log"

	demo "food_delivery/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

//
func NewGRPCStore(db *sqlStore) *grpcStore {
	return &grpcStore{sqlStore: db}
}

type grpcStore struct {
	*sqlStore
	demo.UnimplementedRestaurantLikeServiceServer
}

func (s *grpcStore) GetRestaurantLikeStat(ctx context.Context, request *demo.RestaurantLikeStatRequest) (*demo.RestaurantLikeStatResponse, error) {
	log.Println("gRPC running GetRestaurantLikeState")
	resIds := make([]int, len(request.ResIds))

	for i := range resIds {
		resIds[i] = int(request.ResIds[i])
	}

	mapUserLiked, err := s.GetRestaurantLikes(ctx, resIds)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetRestaurantLikeStat has something error %s", err)
	}

	result := make(map[int32]int32)
	for k, v := range mapUserLiked {
		result[int32(k)] = int32(v)
	}

	return &demo.RestaurantLikeStatResponse{Result: result}, err
}
