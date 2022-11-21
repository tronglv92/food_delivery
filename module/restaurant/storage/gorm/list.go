package restaurantstorage

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	"strings"

	"go.opencensus.io/trace"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	var result []restaurantmodel.Restaurant
	var empty []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("user_id=?", f.OwnerId)
		}
		search := strings.TrimSpace(f.Search)
		if len(search) > 0 {
			db = db.Where("name LIKE ?", "%"+search+"%")
		}
	}

	_, span := trace.StartSpan(context, "store.restaurant.list.count")
	if err := db.Count(&paging.Total).Error; err != nil {
		span.End()
		return empty, err
	}
	span.End()
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return empty, common.ErrDB(err)
		}
		db = db.Where("id<?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	_, span = trace.StartSpan(context, "store.restaurant.list")
	defer span.End()
	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return empty, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
