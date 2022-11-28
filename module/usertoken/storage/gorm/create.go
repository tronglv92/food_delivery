package usertokenstore

import (
	"context"
	"food_delivery/common"
	usertokenmodel "food_delivery/module/usertoken/model"
)

func (s *sqlStore) CreateUserToken(ctx context.Context, data *usertokenmodel.UserTokenCreate) error {
	data.PrepareForInsert()
	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
