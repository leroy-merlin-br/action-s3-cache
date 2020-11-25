package main

import (
	"fmt"
)

func GetOsArgValue(args []string, defaultValue string, argName string) string {
	for i, arg := range args {
		if arg == fmt.Sprintf("--%s", argName) {
			valuePosition := i + 1
			valuePositionLength := valuePosition + 1

			if len(args) >= valuePositionLength {
				return args[valuePosition]
			}
		}
	}

	return defaultValue
}
