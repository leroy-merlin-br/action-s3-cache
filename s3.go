package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
)

// PutObject - Upload object to s3 bucket
func PutObject(key, bucket, s3Class string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load default SDK config: %w", err)
	}
	session := s3.NewFromConfig(cfg)

	file, err := os.Open(key)
	if err != nil {
		return fmt.Errorf("unable to open file %s: %w", key, err)
	}
	defer file.Close()

	i := &s3.PutObjectInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(key),
		Body:         file,
		StorageClass: types.StorageClass(s3Class),
	}

	_, err = session.PutObject(context.TODO(), i)
	if err != nil {
		return fmt.Errorf("unable to put object %s in bucket %s: %w", key, bucket, err)
	}
	
	log.Print("Cache saved successfully")
	return nil
}

// GetObject - Get object from s3 bucket
func GetObject(key, bucket string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load default SDK config: %w", err)
	}
	session := s3.NewFromConfig(cfg)

	i := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	getOutput, err := session.GetObject(context.TODO(), i)
	if err != nil {
		return fmt.Errorf("unable to get object %s from bucket %s: %w", key, bucket, err)
	}

	log.Printf("Cache downloaded successfully, containing %d bytes", getOutput.ContentLength)
	return nil
}

// DeleteObject - Delete object from s3 bucket
func DeleteObject(key, bucket string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load default SDK config: %w", err)
	}
	session := s3.NewFromConfig(cfg)

	i := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err = session.DeleteObject(context.TODO(), i)
	if err != nil {
		return fmt.Errorf("unable to delete object %s from bucket %s: %w", key, bucket, err)
	}
	
	log.Print("Cache purged successfully")
	return nil
}

// ObjectExists - Verify if object exists in s3
func ObjectExists(key, bucket string) (bool, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return false, fmt.Errorf("unable to load default SDK config: %w", err)
	}
	session := s3.NewFromConfig(cfg)

	i := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	if _, err = session.HeadObject(context.TODO(), i); err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return false, nil
		}
		return false, fmt.Errorf("unable to retrieve metadata for object %s on bucket %s: %w", key, bucket, err)
	}

	return true, nil
}
