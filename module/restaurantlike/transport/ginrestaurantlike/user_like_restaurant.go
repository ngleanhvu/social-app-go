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

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.RestaurantLikeCreate{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(db)
		restaurantStore := restaurantstorage.NewSqlStore(db)
		pb := appCtx.GetPubSub()
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, restaurantStore, pb)

		if err := biz.UserLikeRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
