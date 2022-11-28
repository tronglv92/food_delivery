package sessionstore

import (
	"context"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetSession(ctx context.Context, condition map[string]interface{},
	moreKeys ...string) (*sessionmodel.Session, error) {
	var data sessionmodel.Session

	if err := s.db.Table(sessionmodel.Session{}.TableName()).Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
