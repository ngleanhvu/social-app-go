package ginuser

import (
	"crud-go/common"
	"crud-go/component/appctx"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		c.JSON(200, common.SimpleSuccessResponse(u))
	}
}
