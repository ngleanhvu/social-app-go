package ginrestaurant

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/module/restaurant/biz"
	"crud-go/module/restaurant/model"
	"crud-go/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var data restaurantmodule.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(data.Id))
	}
}
