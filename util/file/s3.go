package file

import (
	"bytes"
	"waifu.pics/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3session *session.Session

// InitS3 : Initiate the S3 session, do not open a new session with each upload
func InitS3(Config util.Config) {
	conf := aws.Config{Region: aws.String(Config.S3.REGION), Endpoint: aws.String(Config.S3.ENDPOINT), Credentials: credentials.NewStaticCredentials(Config.S3.ACCESSKEY, Config.S3.SECRETKEY, "")}
	s3session = session.New(&conf)
}

// Upload : Put a file on the S3 container with a buffer
func Upload(buffer bytes.Buffer, mimetype string, filename string, Config util.Config) error {
	uploader := s3manager.NewUploader(s3session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(Config.S3.BUCKET),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buffer.Bytes()),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile : delete a file from the s3 container
func DeleteFile(filename string, Config util.Config) error {
	deleter := s3manager.NewBatchDelete(s3session)

	objects := []s3manager.BatchDeleteObject{{
		Object: &s3.DeleteObjectInput{
			Key:    aws.String(filename),
			Bucket: aws.String(Config.S3.BUCKET),
		},
	}}

	if err := deleter.Delete(aws.BackgroundContext(), &s3manager.DeleteObjectsIterator{Objects: objects}); err != nil {
		return err
	}

	return nil
}
