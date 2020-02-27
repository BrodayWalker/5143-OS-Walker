// +build windows

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["mkdir"] = Mkdir
}

// Mkdir makes a new directory. Currently, the only supported functionality
// includes making a folder in the current directory with a folder name
// supplied by the user as the first argument.
// Example: % mkdir test
// Result: This command would make a new folder called test in the current
// working directory.
func Mkdir(args []string) {
	// Check for a folder name
	if len(args) == 0 {
		fmt.Println("Error: No folder name included.")
	} else {
		// The folder name should be argument 0
		folder := args[0]

		// Make path to folder
		absPath, absErr := filepath.Abs(folder)
		if absErr != nil {
			fmt.Println(absErr)
		}

		// Try to make the folder
		createErr := os.Mkdir(absPath, 0700)
		if createErr != nil {
			fmt.Println(createErr)
		}
	}
}
