package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(action.Bucket),
		Key:    aws.String(action.Key),
		Body:   file,
	})
	if err == nil {
		log.Print("Cache saved successfully")
	}

	return err
}

// GetObject - Get object from s3 bucket
func GetObject(action Action) error {
	session := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(session)

	file, err := os.Create(action.Key)
	if err != nil {
		return err
	}

	size, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: &action.Bucket,
		Key:    &action.Key,
	})

	log.Printf("Cache downloaded successfully, containing %d bytes", size)

	return err
}

// DeleteObject - Delete object from s3 bucket
func DeleteObject(action Action) error {
	session := session.Must(session.NewSession())
	service := s3.New(session)

	_, err := service.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &action.Bucket,
		Key:    &action.Key,
	})
	if err == nil {
		log.Print("Cache purged successfully")
	}

	return err
}

// ObjectExists - Verify if object exists in s3
func ObjectExists(action Action) (bool, error) {
	session := session.Must(session.NewSession())
	service := s3.New(session)

	if _, err := service.HeadObject(&s3.HeadObjectInput{
		Bucket: &action.Bucket,
		Key:    &action.Key,
	}); err != nil {
		if aerr := err.(awserr.Error); aerr.Code() == ErrCodeNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
