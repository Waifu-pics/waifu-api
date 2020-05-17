package util

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3session *session.Session

// InitS3 : Initiate the S3 session, do not open a new session with each upload
func InitS3(Config Config) {
	conf := aws.Config{Region: aws.String(Config.S3.REGION), Endpoint: aws.String(Config.S3.ENDPOINT), Credentials: credentials.NewStaticCredentials(Config.S3.ACCESSKEY, Config.S3.SECRETKEY, "")}
	s3session = session.New(&conf)
}

// Upload : Put a file on the S3 container with a buffer
func Upload(buffer bytes.Buffer, mimetype string, filename string, Config Config) error {
	uploader := s3manager.NewUploader(s3session)

	fmt.Println("Uploading file to S3...")

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(Config.S3.BUCKET),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buffer.Bytes()),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded %s\n", result.Location)
	return nil
}
