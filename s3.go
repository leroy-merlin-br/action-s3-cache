package main

import (
	"fmt"
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

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(action.Bucket),
		Key:    aws.String(action.Key),
		Body:   file,
	})
	if err == nil {
		fmt.Printf("Cache saved successfully at %s", result.Location)
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

	fmt.Printf("%s file downloaded successfully, containing %d bytes", action.Key, size)

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
		fmt.Printf("%s cache purged successfully", action.Key)
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
		if aerr := err.(awserr.Error); aerr.Code() == s3.ErrCodeNoSuchKey {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
