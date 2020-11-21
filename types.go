package main

const (
	PutAction    = "put"
	DeleteAction = "delete"
	GetAction    = "get"
)

type (
	Action struct {
		Action             string
		AwsAccessKeyId     string
		AwsSecretAccessKey string
		Bucket             string
		Key                string
		Artifacts          []string
	}
)
