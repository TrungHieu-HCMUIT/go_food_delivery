package restaurantstorage

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

func (s *sqlStorage) ListDataByConditions(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var results []restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// Optimize offset
	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err != nil {
			// Because we sort desc so we use <
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.
			Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return results, nil
}
