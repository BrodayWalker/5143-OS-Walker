package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/buger/goterm"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["less"] = Less
}

// Less will print out the contents of an input file one page at a time,
// wrapping around words and in between words if the input file line has
// more characters than the terminal has width.
// Usage:
//					less <inputfile>
func Less(args []string) {
	if len(args) < 3 && len(args) > 0 {
		terminalHeight := goterm.Height()
		terminalWidth := goterm.Width()
		fmt.Println("The window size is", terminalHeight, terminalWidth)
		// fmt.Println("Width and Height of terminal:", terminalWidth, terminalHeight)
		input, _ := os.OpenFile(args[0], os.O_RDONLY, 0755)
		scanner := bufio.NewScanner(input)
		lineCounter := 0 // will count number of lines printed to Stdout
		for scanner.Scan() {
			// if we have printed enough lines to file the terminal window, stop printing
			// This is how we print one page at a time
			if lineCounter > terminalHeight {
				lineCounter = 0                       // reset lineCounter for next page
				keyboard := bufio.NewReader(os.Stdin) //create a reader and
				keyboard.ReadString('\n')             // wait for the user to hit enter
			}
			charCount := 0
			line := scanner.Text() // prints number of characters in string 'line'
			for char := range line {
				if charCount < terminalWidth {
					if string(line[char]) == "\t" {
						fmt.Printf("    ")
						charCount += 4 // increase charCount by 4 to accomdate for a tab (4 spaces)
					} else {
						fmt.Printf(string(line[char])) // print the word along with a whitespace after
						charCount++                    // increase charCount by one to include the space from previous instruction
					}
				} else {
					charCount = 0
					lineCounter++

					fmt.Printf("\n") // wrap around, because less doesn't print past the terminal window width
					if string(line[char]) != " " {
						charCount++
						fmt.Printf(string(line[char])) // print the char that would have been the last char
						lineCounter++
					}
				}
			}
			// Mv moves the file given in the first argument to the location of the second
			fmt.Printf("\n")
			lineCounter++
		}
	} else {
		fmt.Println("Error: Incorrect number of arguments/input files.", "Use form 'less inputFile'.")
	}
}
