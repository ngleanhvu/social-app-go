package ginuser

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/component/hasher"
	userbiz "crud-go/module/user/biz"
	usermodel "crud-go/module/user/model"
	userstorage "crud-go/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hasher()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
