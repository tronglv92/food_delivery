package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"food_delivery/common"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access Key ID
			provider.secret, // Secret access key
			""),
	})

	if err != nil {
		log.Fatalln(err)
	}
	provider.session = s3Session
	return provider
}
func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	
	// CLIENT => SERVER => AWS3
	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	// Khi client gửi ảnh lên server, server sẽ úp ảnh lên aws2, vì vậy nó sẽ tốn băng thông gấp đôi:
	// 5M từ client -> Server, 5M từ Server -> AWS3: tổng cộng 10M
	// và client sẽ chờ lâu để nhận kết quả
	// Có thể dùng Presign để cải thiện vấn đề đó
	// Nhưng nhược điểm thì sẽ phức tạp, phải dựng thêm một lambda trên AWS để update lại databse khi client upload file thành công

	// CLIENT => SERVER => TRA VE URL => CLIENT DUNG URL UPLOAD TRUC TIEP LEN AWS3
	// req, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
	// 	Bucket: aws.String(provider.bucketName),
	// 	Key:    aws.String(dst),
	// 	ACL:    aws.String("private"),
	// })
	// req.Presign(time.Second*5) //=> Trả về URL có hiệu lực 5s, và client dùng URL để upload ảnh trực tiếp lên AWS

	if err != nil {
		return nil, err
	}
	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}
	return img, nil
}
