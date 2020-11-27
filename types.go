package main

const (
	// PutAction - Put artifacts
	PutAction    = "put"

	// DeleteAction - Delete artifacts
	DeleteAction = "delete"

	// GetAction - Get artifacts
	GetAction    = "get"
)

type (
	// Action - Input params
	Action struct {
		Action             string
		AwsAccessKeyID     string
		AwsSecretAccessKey string
		Bucket             string
		Key                string
		Artifacts          []string
	}
)
