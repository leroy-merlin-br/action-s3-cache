package main

import (
	"action-s3-cache/modules/archive"
	"fmt"
)

func (action Action) Exec() error {
	switch act := action.Action; act {
	case PutAction:
		_,_ = archive.CreateZip(action.Artifacts, action.Key)
	case GetAction:
		fmt.Print("GetAction")
	case DeleteAction:
		fmt.Print("DeleteAction")
	default: 
		fmt.Printf("Action \"%s\" is not allowed. Valid options are: [%s, %s, %s]", act, PutAction, DeleteAction, GetAction)
	}

	return nil
}
