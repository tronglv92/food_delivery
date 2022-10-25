package ginupload

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/upload/uploadbusiness"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		folder := ctx.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		ctx.JSON(200, common.SimpleSuccessResponse(img))
	}
}
