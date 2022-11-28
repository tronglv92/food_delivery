package sessionstore

import (
	"context"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
)

func (s *sqlStore) CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error {
	data.PrepareForInsert()
	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
