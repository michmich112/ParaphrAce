package s3

import (
	"log"
	"os"
	"server/core/infrastructure"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct {
	bucket   string
	sesh     session.Session
	uploader s3manager.Uploader
}

func New(s *session.Session) infrastructure.StorageClient {
	bucketName := os.Getenv("S3_BUCKET_NAME")
	if bucketName == "" {
		log.Fatalln("[New S3 Storage Client] - Unable to find value for S3_BUCKET_NAME")
	}
	return S3Client{
		bucket:   bucketName,
		sesh:     *s,
		uploader: *s3manager.NewUploader(s),
	}
}

func (s3c S3Client) Save(key string, doc string) (string, error) {
	res, err := s3c.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3c.bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(doc),
	})
	if err != nil {
		log.Printf("[Storage][S3][Error] Unable to upload to Bucket %s with key %s: %v", s3c.bucket, key, err)
		return "", err
	}
	return aws.StringValue(&res.Location), nil
}
