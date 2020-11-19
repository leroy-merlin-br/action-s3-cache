package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("go actionnnn %s", os.Getenv("AWS_REGION"))	
}
