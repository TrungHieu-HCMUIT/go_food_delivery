package restaurantstorage

import (
	"context"
	"errors"
	"fmt"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

func (s *sqlStorage) GetRestaurant(ctx context.Context, id int) (restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	db := s.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Find(&data).Error; err != nil {
		return data, err
	}

	fmt.Println(data)

	if data.Id == 0 {
		return data, errors.New("record not found")
	}

	return data, nil
}
