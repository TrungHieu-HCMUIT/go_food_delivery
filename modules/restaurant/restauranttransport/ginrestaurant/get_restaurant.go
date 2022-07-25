package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go_restaurant/common"
	"go_restaurant/component"
	"go_restaurant/modules/restaurant/restaurantbusiness"
	"go_restaurant/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewGetRestaurantBusiness(store)

		data, err := business.GetRestaurantById(context, id)
		if err != nil {
			panic(err)
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
