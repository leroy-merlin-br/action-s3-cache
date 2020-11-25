package main

import "fmt"

func (action Action) Exec() error {
	switch act := action.Action; act {
	case PutAction:
		fmt.Print("PutAction")
	case GetAction:
		fmt.Print("GetAction")
	case DeleteAction:
		fmt.Print("DeleteAction")
	default: 
		fmt.Printf("Action \"%s\" is not allowed. Valid options are: [%s, %s, %s]", act, PutAction, DeleteAction, GetAction)
	}

	return nil
}
