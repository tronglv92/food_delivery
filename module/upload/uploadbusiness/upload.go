package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/module/upload/uploadmodel"
	"food_delivery/plugin/aws"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	s3       aws.S3
	imgStore CreateImageStorage
}

func NewUploadBiz(s3 aws.S3, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{s3: s3, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	url, err := biz.s3.UploadFileData(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	// img, err := biz.s3.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img := &common.Image{
		Width:     w,
		Height:    h,
		Url:       url,
		Extension: fileExt,
		CloudName: "s3",
	}

	if err := biz.imgStore.CreateImage(ctx, img); err != nil {
		// delete img on S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
