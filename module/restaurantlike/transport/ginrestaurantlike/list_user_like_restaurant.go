package ginrestaurantlike

import (
	"crud-go/common"
	"crud-go/component/appctx"
	restaurantlikebiz "crud-go/module/restaurantlike/biz"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
	restaurantlikestorage "crud-go/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		filter := restaurantlikemodel.Filter{RestaurantId: int(uid.GetLocalID())}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSqlStore(db)
		biz := restaurantlikebiz.NewListUserLikeRestaurantBiz(store)

		data, err := biz.GetUsersLikeRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
