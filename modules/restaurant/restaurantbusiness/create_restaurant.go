package restaurantbusiness

import (
	"context"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStorage interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStorage
}

func NewCreateRestaurantBusiness(store CreateRestaurantStorage) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (biz *createRestaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)
	return err
}
