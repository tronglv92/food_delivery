package usertokenstore

import (
	"context"
	"food_delivery/common"
	usertokenmodel "food_delivery/module/usertoken/model"
)

func (s *sqlStore) GetUserTokens(ctx context.Context, condition map[string]interface{},
	moreKeys ...string) ([]usertokenmodel.UserToken, error) {
	var results []usertokenmodel.UserToken
	var empty []usertokenmodel.UserToken

	if err := s.db.Table(usertokenmodel.UserToken{}.TableName()).Where(condition).Find(&results).Error; err != nil {

		return empty, common.ErrDB(err)
	}

	return results, nil

}
