package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// To run acceptance tests locally pointed at the real AWS, in your local shell:
// $ export S3_TEST_ACC=1
// $ export S3_TEST_BUCKET=some-real-bucket
// Then run:
// $ make test
// Make sure you have an AWS key and secret configured for the tests to use.
const (
	ENV_VAR_TEST_ACC   = "S3_TEST_ACC"
	ENV_VAR_AWS_BUCKET = "S3_TEST_BUCKET"
)

var isAcceptanceTest = os.Getenv(ENV_VAR_TEST_ACC) == "1"

func TestAcceptance(t *testing.T) {
	if !isAcceptanceTest {
		t.Skip(fmt.Sprintf("skipping because %q is not 1", ENV_VAR_TEST_ACC))
	}
	bucket := os.Getenv(ENV_VAR_AWS_BUCKET)
	if bucket == "" {
		t.Fatalf("need %q set to have an S3 bucket for the test", ENV_VAR_AWS_BUCKET)
	}

	// Create a file for testing.
	key := "/tmp/acceptance-tests"
	if err := os.WriteFile(key, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	cacheMgr, err := NewCacheMgr()
	if err != nil {
		t.Fatal(err)
	}

	// Create an object.
	if err := cacheMgr.PutObject(key, bucket, "STANDARD"); err != nil {
		t.Fatal(err)
	}

	// Make sure it exists.
	exists, err := cacheMgr.ObjectExists(key, bucket)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("object was created but doesn't exist")
	}

	// Make sure we can get it.
	if err := cacheMgr.GetObject(key, bucket); err != nil {
		t.Fatal(err)
	}

	// Now make sure we can delete it.
	if err := cacheMgr.DeleteObject(key, bucket); err != nil {
		t.Fatal(err)
	}

	// Make sure it no longer exists.
	exists, err = cacheMgr.ObjectExists(key, bucket)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("object was deleted but still exists")
	}
}

func TestPutObject(t *testing.T) {

	validKey := "/tmp/valid-key"

	// Create a file for testing.
	if err := os.WriteFile(validKey, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	cacheMgr, err := NewCacheMgr()
	if err != nil {
		t.Fatal(err)
	}

	// Here, we overwrite real AWS calls with our mocked ones.
	cacheMgr.Session = &mockedS3Session{
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
			err := cacheMgr.PutObject(testCase.Key, testCase.Bucket, testCase.S3Class)
			if fmt.Sprintf("%+v", err) != testCase.ExpectedErr {
				t.Fatalf("expected %+v but received %+v", testCase.ExpectedErr, err)
			}
		})
	}
}

func TestGetObject(t *testing.T) {

	cacheMgr, err := NewCacheMgr()
	if err != nil {
		t.Fatal(err)
	}

	cacheMgr.Session = &mockedS3Session{
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
			err := cacheMgr.GetObject(testCase.Key, testCase.Bucket)
			if fmt.Sprintf("%+v", err) != testCase.ExpectedErr {
				t.Fatalf("expected %+v but received %+v", testCase.ExpectedErr, err)
			}
		})
	}
}

func TestDeleteObject(t *testing.T) {

	cacheMgr, err := NewCacheMgr()
	if err != nil {
		t.Fatal(err)
	}

	cacheMgr.Session = &mockedS3Session{
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
			err := cacheMgr.DeleteObject(testCase.Key, testCase.Bucket)
			if fmt.Sprintf("%+v", err) != testCase.ExpectedErr {
				t.Fatalf("expected %+v but received %+v", testCase.ExpectedErr, err)
			}
		})
	}
}

func TestObjectExists(t *testing.T) {

	cacheMgr, err := NewCacheMgr()
	if err != nil {
		t.Fatal(err)
	}

	cacheMgr.Session = &mockedS3Session{
		ExpectedKey:    "valid-key",
		ExpectedBucket: "valid-bucket",
	}

	testCases := []struct {
		Name           string
		Key            string
		Bucket         string
		ExpectedResult bool
		ExpectedErr    string
	}{
		{
			Name:           "Valid inputs",
			Key:            "valid-key",
			Bucket:         "valid-bucket",
			ExpectedResult: true,
			ExpectedErr:    "<nil>",
		},
		{
			Name:           "Invalid key",
			Key:            "invalid",
			Bucket:         "valid-bucket",
			ExpectedResult: false,
			ExpectedErr:    "<nil>",
		},
		{
			Name:           "Invalid bucket",
			Key:            "valid-key",
			Bucket:         "invalid",
			ExpectedResult: false,
			ExpectedErr:    "<nil>",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result, err := cacheMgr.ObjectExists(testCase.Key, testCase.Bucket)
			if fmt.Sprintf("%+v", err) != testCase.ExpectedErr {
				t.Fatalf("expected %+v but received %+v", testCase.ExpectedErr, err)
			}
			if result != testCase.ExpectedResult {
				t.Fatalf("expected %t but received %t", testCase.ExpectedResult, result)
			}
		})
	}
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
	if *params.Bucket != m.ExpectedBucket {
		return nil, fmt.Errorf("unexpected bucket: %s", *params.Bucket)
	}
	if *params.Key != m.ExpectedKey {
		return nil, fmt.Errorf("unexpected key: %s", *params.Key)
	}
	return &s3.DeleteObjectOutput{}, nil
}

func (m *mockedS3Session) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	if *params.Bucket != m.ExpectedBucket {
		// This is a real example of the kind of error we get when the bucket doesn't exist.
		return nil, errors.New("operation error S3: PutObject, https response error StatusCode: 404, RequestID: 2GTH7VQBGCFSND7V, HostID: wH0yX/OG80AtRLVQte7zIcKcAqpZa1Dv0g3R5w14gaTdE9xc992aD7aFj+CyK8YI0LFotS7jAbE=, api error NoSuchBucket: The specified bucket does not exist")
	}
	if *params.Key != m.ExpectedKey {
		// This is a real example of the kind of error we get when the bucket exists but the key doesn't.
		return nil, errors.New("operation error S3: HeadObject, https response error StatusCode: 404, RequestID: PCYHBM6JN7JQWC9S, HostID: +Ohjr+/XAy1AXFxs2LxHuweLlWGpxmV3ADJ06olFX/hYX3ln8V9A54iWZCTApKcO/8Q9PdQP0lM=, api error NotFound: Not Found")
	}
	return &s3.HeadObjectOutput{}, nil
}
