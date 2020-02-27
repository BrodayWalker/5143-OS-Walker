// +build !windows

package main

import (
    "fmt"
    "io"
    "os"
)

func init(){
    // Add this command's function to the command mapping
    ComMap["mv"] = Mv
}

func Mv(args []string) {

	// There is a current file (args[0]) and a target name (args[1])
	if len(args) == 2 {
        
		// This just renames some file (if it exists) in the current directory
		// to a new name
		wd, pathErr := os.Getwd()
		if pathErr != nil {
			fmt.Printf("%v\n", pathErr)
		} else {
			oldPath := wd + "/" + args[0]
			newPath := wd + "/" + args[1]

            // Open the original file
            inputFile, openErr := os.Open(oldPath)
            if openErr != nil {
			    fmt.Printf("%v\n", openErr)
            }else{

                // Create a file at the destination path
                outputFile, createErr := os.Create(newPath)
                if createErr != nil {
                    fmt.Printf("%v\n", createErr)
                }else{

                    // Copy the file to the new destination
                    _, copyErr := io.Copy(outputFile, inputFile)
                    if copyErr != nil {
                        fmt.Printf("%v\n", copyErr)
                    }else{

                        // Delete the original
                        delErr := os.Remove(oldPath)
                        if delErr != nil {
                            fmt.Printf("%v\n", delErr)
                        }
                    }
                }
            }
		}
	}
}