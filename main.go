package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	action := Action{
		Action: os.Getenv("ACTION"),
		AwsAccessKeyID: os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Bucket: os.Getenv("BUCKET"),
		Key: fmt.Sprintf("%s.zip", os.Getenv("KEY")),
		Artifacts: strings.Split(strings.TrimSpace(os.Getenv("ARTIFACTS")), "\n"),
	}

	switch act := action.Action; act {
	case PutAction:
		if len(action.Artifacts) <= 0 {
			log.Fatal("no artifacts provided")
		}

		if err := Zip(action); err != nil {
			log.Fatal(err)
		}
	
		if err := PutObject(action); err != nil {
			log.Fatal(err)
		}
	case GetAction:
		if err := GetObject(action); err != nil {
			log.Fatal(err)
		}

		if err := Unzip(action); err != nil {
			log.Fatal(err)
		}
	case DeleteAction:
		fmt.Print("DeleteAction")
	default: 
		fmt.Printf("Action \"%s\" is not allowed. Valid options are: [%s, %s, %s]", act, PutAction, DeleteAction, GetAction)
	}
}
