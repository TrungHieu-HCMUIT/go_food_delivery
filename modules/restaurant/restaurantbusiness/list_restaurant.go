package restaurantbusiness

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/restaurant/restaurantmodel"
)

type ListRestaurantStorage interface {
	ListDataByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	store ListRestaurantStorage
}

func NewListRestaurantBusiness(store ListRestaurantStorage) *listRestaurantBusiness {
	return &listRestaurantBusiness{store: store}
}

func (biz *listRestaurantBusiness) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByConditions(ctx, nil, filter, paging)
	return result, err
}
