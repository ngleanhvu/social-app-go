package ginrestaurant

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/module/restaurant/biz"
	"crud-go/module/restaurant/model"
	"crud-go/module/restaurant/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var data restaurantmodule.RestaurantCreate

		go func() {
			defer common.AppRecover()

			arr := []int{}
			log.Println(arr[0])
		}()

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
