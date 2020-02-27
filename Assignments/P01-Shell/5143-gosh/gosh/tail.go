package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func init() {
	ComMap["tail"] = Tail
}

// Tail prints the final n number of lines of a file.
// By defualt, n is set to 10.
// At some point, build a way to handle being passed an absolute path (split on "\")
// Currently only works if the file is in the current working directory.
func Tail(args []string) {
	// Current working directory only
	if len(args) > 0 {
		// Get path to file
		cwd, cwdErr := os.Getwd()
		if cwdErr != nil {
			fmt.Println("Error getting current working directory")
		}
		// Build absolute path. Arg 0 should be a file name
		abs := cwd + "\\" + args[0]

		// Open the file
		f, fErr := os.Open(abs)
		if fErr != nil {
			log.Fatal(fErr)
		}

		// Scan the file line by line and count the number of lines
		lines := 0
		scanCount := bufio.NewScanner(f)
		for scanCount.Scan() {
			lines++
		}

		// Close the file
		f.Close()
		// Reopen the file
		f, fErr = os.Open(abs)
		if fErr != nil {
			log.Fatal(fErr)
		}

		// Target the last 10 lines by default.
		// Start printing when line number == target line number
		line := 0
		target := lines - 10
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// If there are less than 10 lines, print what is there
			if target > 0 && target < 10 {
				fmt.Println(scanner.Text())
			} else if line >= target {
				fmt.Println(scanner.Text())
			}
			line++
		}

		// Close the file
		f.Close()
	}
}
