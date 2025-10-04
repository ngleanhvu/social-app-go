package ginupload

import (
	"crud-go/common"
	"crud-go/component/appctx"
	uploadbiz "crud-go/module/upload/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDbConnection()
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		defer file.Close()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
