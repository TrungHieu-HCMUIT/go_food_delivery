package ginuser

import (
	"github.com/gin-gonic/gin"
	"go_restaurant/common"
	"go_restaurant/component"
	"go_restaurant/component/hasher"
	"go_restaurant/modules/user/userbusiness"
	"go_restaurant/modules/user/usermodel"
	"go_restaurant/modules/user/userstorage"
	"net/http"
)

func Register(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStorage(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterBusiness(store, md5)

		if err := business.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}

}
