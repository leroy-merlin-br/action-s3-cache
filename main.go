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

		err := Zip(action)
		if err != nil {
			log.Fatal(err)
		}
	
		err = PutObject(action)
		if err != nil {
			log.Fatal(err)
		}
	case GetAction:
		err := GetObject(action)
		if err != nil {
			log.Fatal(err)
		}

		err = Unzip(action)
		if err != nil {
			log.Fatal(err)
		}
	case DeleteAction:
		fmt.Print("DeleteAction")
	default: 
		fmt.Printf("Action \"%s\" is not allowed. Valid options are: [%s, %s, %s]", act, PutAction, DeleteAction, GetAction)
	}
}
