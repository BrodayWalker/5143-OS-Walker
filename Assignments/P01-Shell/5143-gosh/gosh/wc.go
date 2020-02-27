package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["wc"] = Wc
}

//Wc counts the words, lines, and characters of a file
func Wc(args []string) {

	// Split arguments into files and flags
	tokens, flags := ArgSplitter(args)

	// Files object to store content
	files := make(map[string]string)
	for _, fname := range tokens {
		// Read each file
		content, err := ioutil.ReadFile(fname)
		// Check for errors
		if err != nil {
			fmt.Println("Error opening file: ", err)
		}
		// Store the files contents
		files[fname] = string(content)
	}

	// Check which valid flags are in the array of flags
	flagKey := make(map[string]bool)
	for _, flag := range flags {
		flagKey[flag] = true
	}

	// Let's start printing them
	for key, val := range files {
		lines := strconv.Itoa(len(strings.SplitN(val, "\n", -1)))
		words := strconv.Itoa(len(strings.SplitN(val, " ", -1)))
		chars := strconv.Itoa(len(strings.SplitN(val, "", -1)))
		if len(flags) < 1 || flagKey["l"] {
			fmt.Print(lines + " ")
		}
		if len(flags) < 1 || flagKey["w"] {
			fmt.Print(words + " ")
		}
		if len(flags) < 1 || flagKey["m"] {
			fmt.Print(chars + " ")
		}
		fmt.Println(key)
	}
}
