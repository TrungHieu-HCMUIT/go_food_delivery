package restaurantbusiness

import (
	"context"
	"errors"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

type GetRestaurantStorage interface {
	FindDataByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBusiness struct {
	store GetRestaurantStorage
}

func NewGetRestaurantBusiness(storage GetRestaurantStorage) *getRestaurantBusiness {
	return &getRestaurantBusiness{store: storage}
}

func (biz getRestaurantBusiness) GetRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.FindDataByConditions(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if restaurant.Status == 0 {
		return nil, errors.New("data deleted")
	}

	return restaurant, nil
}
