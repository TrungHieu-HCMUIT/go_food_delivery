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
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSqlStorage(ctx.GetMainDBConnection())
		business := restaurantbusiness.NewListRestaurantBusiness(store)

		data, err := business.ListRestaurant(context.Request.Context(), &filter, &paging)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
