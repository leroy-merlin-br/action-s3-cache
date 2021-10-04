package main

import "github.com/aws/aws-sdk-go-v2/service/s3/types"

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
		S3Class   types.StorageClass
		Key       string
		Artifacts []string
	}
)
