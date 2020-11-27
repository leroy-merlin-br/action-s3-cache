package main

import (
	"fmt"
	"log"
)

func (action Action) Exec() error {
	switch act := action.Action; act {
	case PutAction:
		_, err := Zip(action.Artifacts, action.Key)
		if err != nil {
			log.Fatal(err)
			return err
		}
	case GetAction:
		_, err := Unzip(action.Key)
		if err != nil {
			log.Fatal(err)
			return err
		}
	case DeleteAction:
		fmt.Print("DeleteAction")
	default: 
		fmt.Printf("Action \"%s\" is not allowed. Valid options are: [%s, %s, %s]", act, PutAction, DeleteAction, GetAction)
	}

	return nil
}
