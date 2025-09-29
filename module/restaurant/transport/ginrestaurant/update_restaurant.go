package ginrestaurant

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/module/restaurant/biz"
	"crud-go/module/restaurant/model"
	"crud-go/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantmodule.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), &data, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
