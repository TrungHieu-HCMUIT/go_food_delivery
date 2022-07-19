package restaurantstorage

import (
	"context"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

func (s *sqlStorage) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
