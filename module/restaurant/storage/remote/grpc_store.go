package grpcstore

import (
	"context"
	demo "food_delivery/proto"
)

func NewGRPCClient(client demo.RestaurantLikeServiceClient) *grpcClient {
	return &grpcClient{client: client}
}

type grpcClient struct {
	client demo.RestaurantLikeServiceClient
}

func (c *grpcClient) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	resIds := make([]int32, len(ids))

	for i := range resIds {
		resIds[i] = int32(ids[i])
	}
	res, err := c.client.GetRestaurantLikeStat(ctx, &demo.RestaurantLikeStatRequest{ResIds: resIds})
	if err != nil {
		return nil, err
	}

	result := make(map[int]int)
	for k, v := range res.Result {
		result[int(k)] = int(v)
	}
	return result, nil

}
