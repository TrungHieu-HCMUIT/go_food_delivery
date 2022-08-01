package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go_restaurant/common"
	"go_restaurant/component"
	"go_restaurant/modules/restaurant/restaurantbusiness"
	"go_restaurant/modules/restaurant/restaurantmodel"
	"go_restaurant/modules/restaurant/restaurantstorage"
	"net/http"
)

func ListRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filter restaurantmodel.Filter
		if err := context.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantstorage.NewSqlStorage(ctx.GetMainDBConnection())
		business := restaurantbusiness.NewListRestaurantBusiness(store)

		result, err := business.ListRestaurant(context.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
