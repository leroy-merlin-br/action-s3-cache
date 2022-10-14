package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func NewCacheMgr() (*CacheMgr, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return &CacheMgr{
		Session: s3.NewFromConfig(cfg),
	}, nil
}

type CacheMgr struct {
	Session s3Session
}

// PutObject - Upload object to s3 bucket
func (c *CacheMgr) PutObject(key, bucket, s3Class string) error {
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

	if _, err = c.Session.PutObject(context.TODO(), i); err != nil {
		return err
	}
	log.Print("Cache saved successfully")

	return nil
}

// GetObject - Get object from s3 bucket
func (c *CacheMgr) GetObject(key, bucket string) error {
	i := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := c.Session.GetObject(context.TODO(), i)
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
func (c *CacheMgr) DeleteObject(key, bucket string) error {
	i := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	if _, err := c.Session.DeleteObject(context.TODO(), i); err != nil {
		return err
	}
	log.Print("Cache purged successfully")

	return nil
}

// ObjectExists - Verify if object exists in s3
func (c *CacheMgr) ObjectExists(key, bucket string) (bool, error) {
	i := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	if _, err := c.Session.HeadObject(context.TODO(), i); err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Print("Object doesn't exist")
			return false, nil
		}
		return false, err
	}
	log.Print("Object exists")
	return true, nil
}
