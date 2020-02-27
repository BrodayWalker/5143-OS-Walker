package main

import (
	"fmt"
	"regexp"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["echo"] = Echo
}

// Echo repeats what is typed after the command
// Usage:
//				echo statement
func Echo(args []string) {
	// Compile the expression to test against
	// A flag must be in args[0] or it will not be caught
	// TODO: Add support for long-version flags
	var flags, _ = regexp.Compile(`^-{1,2}(\d|\D)`)

	// Does not currently interpret the flag and modify behavior.
	// For now, this ignores the flags in args[0] and prints everything
	// after
	if flags.MatchString(args[0]) {
		// Handle flags
		for i := 1; i < len(args); i++ {
			fmt.Printf(args[i])
		}
		fmt.Printf("\n")
	} else {
		// No flags
		for _, value := range args {
			fmt.Printf(value + " ")
		}
		fmt.Printf("\n")
	}
}
