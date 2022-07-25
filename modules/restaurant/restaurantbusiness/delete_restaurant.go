package restaurantbusiness

import (
	"context"
	"errors"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStorage interface {
	FindDataByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	SoftDelete(ctx context.Context, id int) error
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStorage
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStorage) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (biz *deleteRestaurantBusiness) SoftDeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByConditions(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
