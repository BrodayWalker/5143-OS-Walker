package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["rm"] = Rm
}

// Rm removes a folder or file
func Rm(args []string) {
	if len(args) != 0 {
		argList, flags := ArgSplitter(args)

		// Get absolute path
		path, absErr := filepath.Abs(argList[0])
		if absErr != nil {
			fmt.Println(absErr)
		}

		// Use first non-flag argument as file or folder to delete
		if len(argList) > 0 {
			// Plain delete; must be a file
			if len(flags) == 0 {
				// Try to remove the file
				err := os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
			} else if len(flags) > 0 {
				// Check the flags
				rFlag := false
				fFlag := false
				// Loop through flag array
				for _, v := range flags {
					if v == "r" {
						rFlag = true
					}
					if v == "f" {
						fFlag = true
					}
				}

				// -f is not currently supported so rm -rf <some_folder> works
				// the same as rm -r <some_folder>
				if (rFlag == true && fFlag == true) || (rFlag == true && fFlag == false) {
					os.RemoveAll(path)
				}
			}
		}
	}
}
