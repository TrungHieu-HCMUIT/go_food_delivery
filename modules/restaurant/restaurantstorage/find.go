package restaurantstorage

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStorage) FindDataByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}

	return &result, nil
}
