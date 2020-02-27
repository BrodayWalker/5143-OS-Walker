package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["touch"] = Touch
}

// Touch ...
func Touch(args []string) {
	// If there is an argument
	if len(args) != 0 {
		path := args[0]
		// Make sure argument is a valid file
		exists := FileExists(path)

		// If the file exists, update modification time
		if exists == true {
			// Update access and modification times
			time := time.Now().Local()
			timeErr := os.Chtimes(path, time, time)
			if timeErr != nil {
				fmt.Println(timeErr)
			}
		} else {
			// Attempt to create file
			fp, createErr := os.Create(path)
			if createErr == nil {
				fp.Close()
			} else {
				fmt.Println(createErr)
			}
		}
	}
}
