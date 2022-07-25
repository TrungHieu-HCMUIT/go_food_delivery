package restaurantstorage

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

func (s *sqlStorage) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
