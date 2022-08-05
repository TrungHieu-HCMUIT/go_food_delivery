package restaurantbusiness

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/restaurant/restaurantmodel"
	"log"
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

type ListRestaurantLikeStorage interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBusiness struct {
	store     ListRestaurantStorage
	likeStore ListRestaurantLikeStorage
}

func NewListRestaurantBusiness(store ListRestaurantStorage, likeStore ListRestaurantLikeStorage) *listRestaurantBusiness {
	return &listRestaurantBusiness{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBusiness) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByConditions(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println("Cannot get list restaurant: ", err)
	}

	if m := mapResLike; m != nil {
		for i, restaurant := range result {
			result[i].LikedCount = mapResLike[restaurant.Id]
		}
	}

	return result, nil
}
