package util

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3session *session.Session

func InitS3(Config Config) {
	conf := aws.Config{Region: aws.String(Config.S3.REGION), Endpoint: aws.String(Config.S3.ENDPOINT), Credentials: credentials.NewStaticCredentials(Config.S3.ACCESSKEY, Config.S3.SECRETKEY, "")}
	s3session = session.New(&conf)
}

func Upload(buffer bytes.Buffer, mimetype string, filename string, Config Config) {
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
		fmt.Println("error", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully uploaded %s\n", result.Location)
}
