package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	// scanner pointer that will point to a file to be read from
	scanner *bufio.Scanner

	// this string will save the path of the file to which output will be redirected when > or >> are used
	// and ctext will hold strings typed into standard in when 'cat' is called with no arguments
	ctext string
)

func init() {
	// Add this command's function to the command mapping
	ComMap["cat"] = Cat
}

// Cat will either concatenate a file and print it to std out or
// intake two or more files and print them as if they were concatenated.
// If Cat was called without any arguments, it will take input from Stdin
// and just output it to Stdout.
// Usage:
//			cat
// 			cat file
// 			cat file1 file2 fileN
func Cat(args []string) {
	// base case, where there are no arguments to the command.
	// Read from Stdin and output straight to Stdout until a
	// 'q' is typed
	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for ctext != "q\r\n" {
			ctext, _ = reader.ReadString('\n')
			fmt.Println(ctext)
		}
		//if cat has arguments, create a temporary file. We will use this to
		//store contents of the input files.
	} else {
		tempFile, err := os.OpenFile("temp.txt", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			log.Fatal(err)
		}
		for element := range args {
			file, _ := os.Open(args[element]) //open the file
			scanner = bufio.NewScanner(file)  //create a scanner to run through the file
			scanner.Split(bufio.ScanLines)    //set delimiter to lines, meaning the scanner will read in one line at a time
			for scanner.Scan() {              //run through the file and print each line to the stdout
				tempFile.WriteString(scanner.Text() + "\n") //write to the temp file
			}
			file.Close()
		}
		tempFile.Seek(0, io.SeekStart) //return scanner to the top of tempFile
		scanner = bufio.NewScanner(tempFile)
		scanner.Split(bufio.ScanLines)
		// print the contents of tempFile to stdout
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
