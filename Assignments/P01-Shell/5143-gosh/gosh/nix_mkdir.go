// +build !windows

package main

import (
	"fmt"
	"os"
	"syscall"
)

func init(){
    // Add this command's function to the command mapping
    ComMap["mkdir"] = Mkdir
}

// Mkdir...
func Mkdir(args []string) {
	// Get the current working directory
	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Failed to get current working directory.")
	}

	// Check for a folder name
	if len(args) == 0 {
		fmt.Println("Error: No folder name included.")
	} else {
		// The folder name should be argument 0
		folderName := args[0]
		// Construct absolute path
		totalPath := path + "/" + folderName

		// Print path for testing
		fmt.Println(totalPath)

		// Make the directory using the Mkdir system call (no lower level option)
		errPath := syscall.Mkdir(totalPath, 0755)

		if errPath == syscall.EEXIST {
			fmt.Println("Directory already exists.")
		} else if errPath == syscall.ENOENT {
			fmt.Println("Failed to create directory. One or more intermediate directories do not exist.")
		}
	}
}
