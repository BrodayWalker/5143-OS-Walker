// +build windows

package main

import (
	"fmt"
	"os"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["mv"] = Mv
}

// Mv moves the file given in the first argument to the location of the second argument.
// Mv also renames files.
func Mv(args []string) {
	// There is a current file (args[0]) and a target name (args[1])
	if len(args) == 2 {
		oldPath := args[0]
		newPath := args[1]
		// Find out if our files exist
		oldPathIsFile := FileExists(oldPath)
		oldPathIsDir := ValidPathToDir(oldPath)

		// If the old path exists (this command does nothing if it does not exist)
		if oldPathIsFile == true || oldPathIsDir == true {
			// If old path is a file, it can be moved into a directory or renamed
			if oldPathIsFile == true {
				oldPath = BuildPathToFile(oldPath)
			} else if oldPathIsDir == true {
				oldPath = BuildPathToDir(oldPath)
			}

			renameErr := os.Rename(oldPath, newPath)

			if renameErr != nil {
				fmt.Println(renameErr)
			}
		}
	}
}
