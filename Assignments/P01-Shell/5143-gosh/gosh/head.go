package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func init(){
    // Add this command's function to the command mapping
    ComMap["head"] = Head
}

// Head prints the first 10 lines of a file or the number of lines
// specified after the -n flag
func Head(args []string) {
	// Print test file path
	if len(args) != 0 {
		filePath := args[0]
		fmt.Println(filePath)
		// Print full path for testing
		fullPath, ferr := os.Getwd()
		if ferr != nil {
			fmt.Println(ferr)
		} else {
			fullPath = fullPath + "/" + filePath
			fmt.Println(fullPath)
		}

		// Number of lines to print. Print 10 by default or all lines that exist
		// if the file is less than 10 lines
		n := 10
		if len(args) > 1 {
			if args[1] == "-n" {
				n, _ = strconv.Atoi(args[2])
			}
		}

		// Open the file
		file, err := os.Open(fullPath)
		// Print errors, if any
		if err != nil {
			fmt.Println(err)
		}

		// A reader object is used to read the file
		reader := bufio.NewReader(file)

		// Read n lines and print them
		for i := 0; i < n; i++ {
			line, _ := reader.ReadString('\n')
			fmt.Printf(line)
		}
		fmt.Printf("\n")

		file.Close()
	} else {
		fmt.Println("Not enough arguments")
	}
}
