package amazon

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

type S3Service struct {
	Session *SessionAWS
	Bucket  string
}

func NewS3Service(credential *SessionAWS, bucketName string) *S3Service {
	return &S3Service{
		Session: credential,
		Bucket:  bucketName,
	}
}

func (service *S3Service) Upload(fileName string, file io.Reader) error {
	uploader := s3manager.NewUploader(service.Session.Config)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &service.Bucket, // バケット名を指定してください
		Key:    &fileName,       // S3上のファイル名を指定してください
		Body:   file,
	})
	if err != nil {
		return err
	}
	fmt.Println("Successfully uploaded to", "myBucket")
	return nil
}

type UploaderI interface {
	Upload(fileName string, file io.Reader) error
}

func S3Client() (UploaderI, error) {
	sessionAws, err := NewSessionAWS()
	if err != nil {
		return nil, err
	}

	s3Service := NewS3Service(sessionAws, os.Getenv("AWS_BUCKET"))
	return s3Service, nil
}
