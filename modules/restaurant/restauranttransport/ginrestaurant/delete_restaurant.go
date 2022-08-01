package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go_restaurant/common"
	"go_restaurant/component"
	"go_restaurant/modules/restaurant/restaurantbusiness"
	"go_restaurant/modules/restaurant/restaurantstorage"
	"net/http"
)

func DeleteRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		//id, err := strconv.Atoi(context.Param("id"))
		uid, err := common.FromBase58(context.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStorage(ctx.GetMainDBConnection())
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err := business.SoftDeleteRestaurant(context.Request.Context(), int(uid.GetLocalID())); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
