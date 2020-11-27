package main
import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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
func GetObject() error{


	return nil
}

// DeleteObject - Delete object from s3 bucket
func DeleteObject() {

}