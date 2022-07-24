package restaurantstorage

import (
	"context"
	"go_restaurant/modules/restaurant/restaurantmodel"
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

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *sqlStorage) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
