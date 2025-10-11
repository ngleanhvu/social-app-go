package ginrestaurantlike

import (
	"crud-go/common"
	"crud-go/component/appctx"
	restaurantstorage "crud-go/module/restaurant/storage"
	restaurantlikebiz "crud-go/module/restaurantlike/biz"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
	restaurantlikestorage "crud-go/module/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.RestaurantLikeUpdate{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(db)
		restaurantStore := restaurantstorage.NewSqlStore(db)
		biz := restaurantlikebiz.NewDislikeRestaurantBiz(store, restaurantStore, restaurantStore)

		if err := biz.UserDislikeRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
