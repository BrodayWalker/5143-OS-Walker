package main

import (
	"fmt"
	"io"
	"os"
)

// PipeLine :
// Takes a list of Commands, then for each Command, sets it's output to a
// file and uses that file as the first arg of the next command, executing
// them all in sequence.
// Notes:
//   - Actual command calling in handled in Execute()
//   - In the event of an input redirection in use with Pipes, redirect is
//     handled by Execute(), which will in turn call RedirectAndExecute()
func PipeLine(commands []Command) {

	// stdout backup
	stdout := os.Stdout
	// error variable
	var err error

	// Path to the pipe file
	dir := os.TempDir()
	outPipeName := "pipe.out.tmp"
	if err != nil {
		fmt.Println("Error getting working directory: ", err)
		return
	}
	outPipePath := dir + "/" + outPipeName
	inPipeName := "pipe.tmp"
	if err != nil {
		fmt.Println("Error getting working directory: ", err)
		return
	}
	inPipePath := dir + "/" + inPipeName

	// For each command in the array
	for i, pipe := range commands {

		// If this isn't the last command
		if i < len(commands)-1 {
			// Before processing each command, open the file and redirect stdout
			os.Stdout, err = os.OpenFile(outPipePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				os.Stdout = stdout
				fmt.Println("Error opening temp file: ", err)
				return
			}
		}

		// If this isn't the first command
		if i > 0 {
			// Add the pipe file to the args (at the front)
			pipe.args = frAddStr(pipe.args, inPipePath)
		}
		// Execute the command with its arguments
		Execute(pipe)

		// After processing each command, if this isn't the last command
		if i < len(commands)-1 {
			// Close the pipe file
			err = os.Stdout.Close()
			if err != nil {
				fmt.Println("Error closing temp file: ", err)
				return
			}
		}
		// After processing each command, restore stdout
		os.Stdout = stdout
		// Copy the output to the input file for the next command
		copyFrom(outPipePath, inPipePath)
	}
	// Remove the temporary files
	Rm([]string{inPipePath})
	Rm([]string{outPipePath})
}

// Add an item to the front of an array of strings
func frAddStr(argList []string, arg string) []string {
	argList = append(argList, "")
	copy(argList[1:], argList)
	argList[0] = arg
	return argList
}

// Copy a file from src to dest
// copy() borrowed from
// https://opensource.com/article/18/6/copying-files-go
func copyFrom(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
