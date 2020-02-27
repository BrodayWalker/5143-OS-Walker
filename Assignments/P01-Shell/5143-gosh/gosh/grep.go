package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["grep"] = Grep
}

//Grep ..
func Grep(args []string) {
	// Merge any quoted arguments
	args = MergeQuotedArgs(args)
	// Split arguments into regexp, files and flags
	tokens, flags := ArgSplitter(args)
	expr := tokens[0]
	files := tokens[1:]
	// Trim the quotes as Go's re library doesn't use them
	expr = strings.TrimLeft(expr, "\"")
	expr = strings.TrimRight(expr, "\"")

	// Check which valid flags are in the array of flags
	flagKey := make(map[string]bool)
	for _, flag := range flags {
		flagKey[flag] = true
	}

	// Compile the regular expression
	re, err := regexp.Compile(expr)
	if err != nil {
		fmt.Println("Error compiling regular expression: ", err)
	}

	// For every file argument sent
	for _, file := range files {
		// Get the matches for that file
		matches := _grep(re, file)
		// If the -l flag is set, just print filenames of files that have matches
		if flagKey["l"] {
			if len(matches) > 0 {
				fmt.Println(file)
			}
		} else {
			// If no -l flag, print every matched line
			for _, match := range matches {
				fmt.Println(match)
			}
		}
	}

}

func _grep(re *regexp.Regexp, filename string) []string {
	// Open the file and start reader
	file, err := os.Open(filename)
	freader := bufio.NewReader(file)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()
	// Create a string and error variable
	var line string
	var readErr error
	// List of matches
	matches := make([]string, 0)
	// Loop through each line of the file
	for readErr != io.EOF {
		line, readErr = freader.ReadString('\n')
		// Don't alert for an EOF error, just let the loop finish
		if readErr != nil && readErr != io.EOF {
			fmt.Println("Error opening file: ", err)
		}
		// Find all the matches
		if re.MatchString(line) {
			matches = append(matches, line)
		}
	}
	return matches
}
