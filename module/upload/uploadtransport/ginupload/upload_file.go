package ginupload

import (
	"food_delivery/common"
	"food_delivery/module/upload/uploadbusiness"
	"food_delivery/module/upload/uploadstorage"
	"food_delivery/plugin/aws"
	goservice "food_delivery/plugin/go-sdk"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upload(sc goservice.ServiceContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		s3 := sc.MustGet("aws").(aws.S3)
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbusiness.NewUploadBiz(s3, imgStore)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
