package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["cp"] = Cp
}

// Cp will copy a file or a set of files into another location.
// This function only work if destinations are Unix-formatted, i.e.
// path/to/destination, NOT path\to\destination OR path\\to\\destination
//Usage:
//			cp path/to/file path/to/copy/to
//			cp path/to/file1 path/to/file2 path/to/fileN path/to/copy/to
func Cp(args []string) {
	if len(args) > 1 {
		// args[lastIndex] is the destination folder for the file(s) to be copied
		lastIndex := len(args) - 1
		// check if destination exists. If it doesn't, exit command
		if _, err := os.Stat(args[lastIndex]); !os.IsNotExist(err) || filepath.Ext(args[lastIndex]) != "" {
			for file := 0; file < lastIndex; file++ {
				// check if the source files exist, and if they don't skip and move on
				// to the next file to copy
				if _, err := os.Stat(args[file]); !os.IsNotExist(err) {
					// open the file to be copied for reading only
					copyFrom, errFrom := os.OpenFile(args[file], os.O_RDONLY, 0755)
					if errFrom != nil {
						fmt.Println(err)
					}
					// open a file of the same name in the destination folder
					copyTo, errTo := os.OpenFile(args[lastIndex], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
					if err != nil {
						fmt.Println(errTo)
					}
					// create a scanner to to read through the source file
					scanner := bufio.NewScanner(copyFrom)
					// copy source file to destination one line at a time
					for scanner.Scan() {
						copyTo.WriteString(string(scanner.Text()) + "\n")
					}
					copyTo.Close()
					copyFrom.Close()
				} else {
					fmt.Println("Error: " + args[file] + " does not exist. Skipping.")
				}
			}
		} else {
			fmt.Println("Error: Destination folder does not exist.")
		}
	} else {
		fmt.Println("Error: Too few arguments. Must have a file to copy and a destination.")
	}
}
