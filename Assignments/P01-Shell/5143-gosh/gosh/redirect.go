package main

import (
	"fmt"
	"os"
)

// RedirectAndExecute :
// Special version of the execute function which takes a command and runs it,
// directing its input and output to files
func RedirectAndExecute(command Command) {

	// Backups
	stdin := os.Stdin
	stdout := os.Stdout
	// error variable
	var err error

	// Redirect the in/out streams
	if command.infile != nil {
		os.Stdin, err = os.OpenFile(*command.infile, os.O_RDWR, 0755)
	}
	if command.outfile != nil {
		if command.appendMode {
			os.Stdout, err = os.OpenFile(*command.outfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		} else {
			os.Stdout, err = os.OpenFile(*command.outfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		}
	}
	if err != nil {
		os.Stdout = stdout
		fmt.Println("Error opening file: ", err)
		return
	}

	// If the Command is valid
	if com, valid := ComMap[command.key]; valid {
		// Execute the command with its arguments
		com(command.args)
	}

	// If input was redirected, close the file when we're done
	if os.Stdin != stdin {
		err = os.Stdin.Close()
	}
	// Then reset stdin
	os.Stdin = stdin
	// If output was redirected, close the file when we're done
	if os.Stdout != stdout {
		err = os.Stdout.Close()
	}
	// Then reset stdin
	os.Stdout = stdout
	// Check for errors
	if err != nil {
		fmt.Println("Error closing file: ", err)
		return
	}
}
