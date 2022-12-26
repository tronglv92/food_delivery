package userstore

import (
	"context"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
)

func (s *sqlStore) CreateUserFacebook(ctx context.Context, data *usermodel.UserFacebookCreate) error {
	data.PrepareForInsert()
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
