package main

type (
	Action struct {
		Action             string
		AwsAccessKeyId     string
		AwsSecretAccessKey string
		Bucket             string
		Key                string
		Artifacts          string
	}
)
