package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// Exclamation will return the n-th command from the commandList
// and execute it automatically. If the command does not exist,
// it returns an error message and the shell continues like normal.
// In this function, location is the command number that follows the
// '!' character
// Usage:
//			!43 || !3 || !<whatever number you want>
func Exclamation(location string) {
	//convert the location string to an integer
	element, _ := strconv.Atoi(location)
	//check if the command actually exists in commandList,
	//and if it does, find it
	if element < len(commandList) {
		checker = true
		singleCommand := commandList[element]
		//open a temp file to store the command. If the temp file exists, truncate it
		tmpfile, _ = os.OpenFile("temp.tmp", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
		tmpfile.WriteString(singleCommand + "\n")
		tmpfile.Seek(0, io.SeekStart)
		fmt.Println(commandList[element])
	} else {
		//return an error if the command doesn't exist. Shell resumes functionality normally
		fmt.Println("Error: command", location, "does not exist!")
	}
}
