package main
import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// PutObject - Upload object to s3 bucket
func PutObject(action Action) error {
	session := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(session)

	file, err := os.Open(action.Key)
	if err != nil {
		return err
	}
	defer file.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(action.Bucket),
		Key: aws.String(action.Key),
		Body: file,
	})

	if err == nil {
		fmt.Printf("Cache successfully saved at %s", result.Location)
	}

	return err
}

// GetObject - Get object from s3 bucket
func GetObject(action Action) error {
	session := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(session)

	file, err := os.Create(action.Key)
	if err != nil {
		return nil
	}

	size, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: &action.Bucket,
		Key: &action.Key,
	})
	
	if err == nil {
		fmt.Printf("%s file downloaded with %d bytes", action.Key, size)
	}

	return err
}

// DeleteObject - Delete object from s3 bucket
func DeleteObject() {

}