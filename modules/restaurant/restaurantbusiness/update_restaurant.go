package restaurantbusiness

import (
	"context"
	"errors"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStorage interface {
	FindDataByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStorage
}

func NewUpdateRestaurantBusiness(store UpdateRestaurantStorage) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (biz *updateRestaurantBusiness) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataByConditions(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
