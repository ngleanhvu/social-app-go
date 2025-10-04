package ginuser

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/component/hasher"
	"crud-go/component/tokenprovider/jwt"
	userbiz "crud-go/module/user/biz"
	usermodel "crud-go/module/user/model"
	userstorage "crud-go/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var userLogin usermodel.UserLogin

		if err := c.ShouldBind(&userLogin); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hasher()
		tokenprovider := jwt.NewJWTProvider(common.SecretKey)
		biz := userbiz.NewLoginBiz(store, tokenprovider, md5, 1000000000)

		token, err := biz.Login(c.Request.Context(), &userLogin)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
