package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Session interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
}

const (
	// PutAction - Put artifacts
	PutAction = "put"

	// DeleteAction - Delete artifacts
	DeleteAction = "delete"

	// GetAction - Get artifacts
	GetAction = "get"

	// ErrCodeNotFound - s3 Not found error code
	ErrCodeNotFound = "NotFound"
)

type (
	// Action - Input params
	Action struct {
		Action    string
		Bucket    string
		S3Class   string
		Key       string
		Artifacts []string
	}
)
