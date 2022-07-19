package restaurantbusiness

import (
	"context"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

type GetRestaurantStorage interface {
	GetRestaurant(ctx context.Context, id int) (restaurantmodel.Restaurant, error)
}

type getRestaurantBusiness struct {
	store GetRestaurantStorage
}

func NewGetRestaurantBusiness(storage GetRestaurantStorage) *getRestaurantBusiness {
	return &getRestaurantBusiness{store: storage}
}

func (biz getRestaurantBusiness) GetRestaurantById(ctx context.Context, id int) (restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.GetRestaurant(ctx, id)
	if err != nil {
		return restaurantmodel.Restaurant{}, err
	}

	return restaurant, nil
}
