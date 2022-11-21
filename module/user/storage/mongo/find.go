package userstore

import (
	"context"
	usermodel "food_delivery/module/user/model"
)

func (s *mgoStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	return nil, nil
}
