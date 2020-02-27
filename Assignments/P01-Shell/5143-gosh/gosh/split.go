package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["split"] = Split
}

var lines, size, chunks, length, counter int = 0, 0, 0, 0, 0
var file, file2 *os.File
var err error

// Split will divide an input file into a certain number of
// individual files depending on arguments. If no arguments are sent in
// the file is divided into 10 parts (well, a bug makes it divide into 11 parts,
// but the 11th file is a duplicate of the 10th. If -l <n> is used,
// the file is divided into parts with n number of lines each.
// Usage:
//		split [option](-l) [option value](<n>) [input](infile)
func Split(args []string) {
	//loop will iterate through all elements in args, which will will be flags, values, and/or input files
	for word := 0; word < len(args); word++ {
		//switch cases will only cover single-dashed flags for now
		switch string(args[word][0]) {
		//if there is a flag, it's going to be an l flag
		case "-":
			//grabbing the number of lines of text to put in each child file
			lines, _ = strconv.Atoi(args[word+1])
			//slice off the processed flag and associated lines value
			args = args[1:]
		default:
			//if the lines flag was waved
			if lines != 0 {
				printByLines(lines, args[word])
				args = args[word:]
			} else {
				counter = 0
				file, err = os.OpenFile(args[word], os.O_RDONLY, 0755) //open file from input string array
				if err != nil {
					log.Fatal(err)
				}
				fi, _ := file.Stat()
				//this variable will hold the sections of data we will be writing to separate files
				data := make([]byte, fi.Size()/10)
				for counter < 10 {
					//create a "sub-file" with the same name as the parent but with a counter value appended to the front
					file2, err = os.OpenFile(strconv.Itoa(counter)+filepath.Base(args[word]), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
					if err != nil {
						log.Fatal(err)
					}
					//read the data from 'file' into array 'data'
					_, err = file.Read(data)
					//write contents of 'data' to 'file2'
					file2.Write(data)
					file2.Close()
					counter++
				}
				file.Close()
			}
		}
	}
}

// function that will split file on 'path' into
// separate files of 'lineNum' lines each
func printByLines(lineNum int, path string) {
	var counter, lineCounter int = 0, 0 //counter: used to name child files; lineCounter: used to count 100 lines per child file
	file, err = os.Open(path)           //open file from input string array
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) //will traverse file on 'path'
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() { //read in a line until eof. Putting a Scan() here will ensure the loop ends at eof
		//create a "sub-file" with the same name as the parent but with a counter value appended to the front
		file2, err = os.Create(strconv.Itoa(counter) + filepath.Base(path))
		if err != nil {
			log.Fatal(err)
		}
		// write the first line returned from the Scan() and let the following loop grab the rest
		file2.WriteString(scanner.Text() + "\n")
		lineCounter = 0
		for lineCounter < lineNum-1 && scanner.Scan() { //Scan() "grabs" a line from the parent file
			file2.WriteString(scanner.Text() + "\n") //Text() returns the line from Scan() as type string
			lineCounter++
		}
		counter++
	}
}
