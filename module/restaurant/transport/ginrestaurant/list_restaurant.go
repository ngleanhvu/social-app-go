package ginrestaurant

import (
	"crud-go/common"
	"crud-go/component/appctx"
	restaurantbiz "crud-go/module/restaurant/biz"
	restaurantmodule "crud-go/module/restaurant/model"
	restaurantstorage "crud-go/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBindQuery(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		pagingData.Fulfill()

		var filter restaurantmodule.Filter

		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		filter.Status = []int{1}

		var result []restaurantmodule.Restaurant

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurantBiz(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
			return
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(200, common.NewSuccessResponse(result, filter, pagingData))
	}
}
