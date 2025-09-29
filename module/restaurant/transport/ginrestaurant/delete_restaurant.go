package ginrestaurant

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/module/restaurant/biz"
	"crud-go/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("uid"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
