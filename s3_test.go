package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestPutObject(t *testing.T) {

	validKey := "/tmp/valid-key"

	// Create a file for testing.
	if err := os.WriteFile(validKey, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	// Here, we overwrite real AWS calls with our mocked ones.
	session = &mockedS3Session{
		ExpectedKey:    validKey,
		ExpectedBucket: "valid-bucket",
	}

	testCases := []struct {
		Name        string
		Key         string
		Bucket      string
		S3Class     string
		ExpectedErr string
	}{
		{
			Name:        "Valid inputs",
			Key:         validKey,
			Bucket:      "valid-bucket",
			S3Class:     "valid-s3class",
			ExpectedErr: "<nil>",
		},
		{
			Name:        "Invalid key",
			Key:         "invalid",
			Bucket:      "valid-bucket",
			S3Class:     "valid-s3class",
			ExpectedErr: "open invalid: no such file or directory",
		},
		{
			Name:        "Invalid bucket",
			Key:         validKey,
			Bucket:      "invalid",
			S3Class:     "valid-s3class",
			ExpectedErr: "unexpected bucket: invalid",
		},
		{
			Name:        "Invalid s3class",
			Key:         validKey,
			Bucket:      "valid-bucket",
			S3Class:     "invalid",
			ExpectedErr: "<nil>",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := PutObject(testCase.Key, testCase.Bucket, testCase.S3Class)
			if fmt.Sprintf("%+v", err) != testCase.ExpectedErr {
				t.Fatalf("expected %+v but received %+v", testCase.ExpectedErr, err)
			}
		})
	}
}

func TestGetObject(t *testing.T) {

	session = &mockedS3Session{
		ExpectedKey:    "valid-key",
		ExpectedBucket: "valid-bucket",
	}

	testCases := []struct {
		Name        string
		Key         string
		Bucket      string
		ExpectedErr string
	}{
		{
			Name:        "Valid inputs",
			Key:         "valid-key",
			Bucket:      "valid-bucket",
			ExpectedErr: "<nil>",
		},
		{
			Name:        "Invalid key",
			Key:         "invalid",
			Bucket:      "valid-bucket",
			ExpectedErr: "unexpected key: invalid",
		},
		{
			Name:        "Invalid bucket",
			Key:         "valid-key",
			Bucket:      "invalid",
			ExpectedErr: "unexpected bucket: invalid",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := GetObject(testCase.Key, testCase.Bucket)
			if fmt.Sprintf("%+v", err) != testCase.ExpectedErr {
				t.Fatalf("expected %+v but received %+v", testCase.ExpectedErr, err)
			}
		})
	}
}

func TestDeleteObject(t *testing.T) {
	// TODO
}

func TestObjectExists(t *testing.T) {
	// TODO
}

type mockedS3Session struct {
	ExpectedKey    string
	ExpectedBucket string
}

func (m *mockedS3Session) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	if *params.Bucket != m.ExpectedBucket {
		return nil, fmt.Errorf("unexpected bucket: %s", *params.Bucket)
	}
	if *params.Key != m.ExpectedKey {
		return nil, fmt.Errorf("unexpected key: %s", *params.Key)
	}
	return &s3.PutObjectOutput{}, nil
}

func (m *mockedS3Session) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	if *params.Bucket != m.ExpectedBucket {
		return nil, fmt.Errorf("unexpected bucket: %s", *params.Bucket)
	}
	if *params.Key != m.ExpectedKey {
		return nil, fmt.Errorf("unexpected key: %s", *params.Key)
	}
	return &s3.GetObjectOutput{}, nil
}

func (m *mockedS3Session) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	// TODO
	return nil, nil
}

func (m *mockedS3Session) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	// TODO
	return nil, nil
}
