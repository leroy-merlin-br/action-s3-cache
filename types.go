package main

const (
	PutAction    = "put"
	DeleteAction = "delete"
	GetAction    = "get"
)

type (
	Action struct {
		Action             string
		AwsAccessKeyID     string
		AwsSecretAccessKey string
		Bucket             string
		Key                string
		Artifacts          []string
	}
)
