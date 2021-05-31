package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	action := Action{
		Action:    os.Getenv("ACTION"),
		Bucket:    os.Getenv("BUCKET"),
		S3Class:   os.Getenv("S3_CLASS"),
		Key:       fmt.Sprintf("%s.zip", os.Getenv("KEY")),
		Artifacts: strings.Split(strings.TrimSpace(os.Getenv("ARTIFACTS")), "\n"),
	}

	switch act := action.Action; act {
	case PutAction:
		if len(action.Artifacts[0]) <= 0 {
			log.Fatal("No artifacts patterns provided")
		}

		if err := Zip(action.Key, action.Artifacts); err != nil {
			log.Fatal(err)
		}

		if err := PutObject(action.Key, action.Bucket, action.S3Class); err != nil {
			log.Fatal(err)
		}
	case GetAction:
		exists, err := ObjectExists(action.Key, action.Bucket)
		if err != nil {
			log.Fatal(err)
		}

		// Get and and unzip if object exists
		if exists {
			if err := GetObject(action.Key, action.Bucket); err != nil {
				log.Fatal(err)
			}

			if err := Unzip(action.Key); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Printf("No caches found for the following key: %s", action.Key)
		}
	case DeleteAction:
		if err := DeleteObject(action.Key, action.Bucket); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Action \"%s\" is not allowed. Valid options are: [%s, %s, %s]", act, PutAction, DeleteAction, GetAction)
	}
}
