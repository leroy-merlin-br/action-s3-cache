package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
)

var session = func() s3Session {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return s3.NewFromConfig(cfg)
}()

// PutObject - Upload object to s3 bucket
func PutObject(key, bucket, s3Class string) error {
	file, err := os.Open(key)
	if err != nil {
		return err
	}
	defer file.Close()

	i := &s3.PutObjectInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(key),
		Body:         file,
		StorageClass: types.StorageClass(s3Class),
	}

	_, err = session.PutObject(context.TODO(), i)
	if err == nil {
		log.Print("Cache saved successfully")
	}

	return err
}

// GetObject - Get object from s3 bucket
func GetObject(key, bucket string) error {
	i := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := session.GetObject(context.TODO(), i)
	if err != nil {
		return err
	}
	log.Printf("Cache downloaded successfully")
	if output.ContentRange != nil {
		log.Printf("Content range is %s", *output.ContentRange)
	}
	return nil
}

// DeleteObject - Delete object from s3 bucket
func DeleteObject(key, bucket string) error {
	i := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := session.DeleteObject(context.TODO(), i)
	if err == nil {
		log.Print("Cache purged successfully")
	}

	return err
}

// ObjectExists - Verify if object exists in s3
func ObjectExists(key, bucket string) (bool, error) {
	i := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	if _, err := session.HeadObject(context.TODO(), i); err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return false, nil
		}
	}

	return true, nil
}
