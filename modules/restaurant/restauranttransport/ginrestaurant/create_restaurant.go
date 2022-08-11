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

func CreateRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := context.MustGet(common.CurrentUser).(common.Requester)
		data.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewSqlStorage(ctx.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)

		if err := business.CreateRestaurant(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
