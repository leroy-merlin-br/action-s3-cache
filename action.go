package main

import (
	"fmt"
)

func (action Action) Exec() error {
	fmt.Print("go action :) \n")

	fmt.Printf("action: %s \n", action.Action)
	fmt.Printf("aws-secret-access-key: %s \n", action.AwsSecretAccessKey)
	fmt.Printf("aws-access-key-id: %s \n", action.AwsAccessKeyId)
	fmt.Printf("bucket: %s \n", action.Bucket)
	fmt.Printf("key: %s \n", action.Key)
	fmt.Printf("artifacts: %s \n", action.Artifacts)
	fmt.Printf("artifacts len: %d \n", len(action.Artifacts))

	return nil
}
