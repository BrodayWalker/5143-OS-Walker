package main

import "fmt"

func init() {
	// Add this command's function to the command mapping
	ComMap["history"] = History
}

// History will print off all commands and their order run
// since the beginning of the shell's creation, including
// commands from previous sessions.
// Usage:
//		history
// Since history doesn't take any arguments, args is useless,
// but leaving it as a parameter allows the shell to still run
// even if the programmer adds arguments for whatever reason
func History(args []string) {
	for i := range commandList {
		fmt.Println(i, commandList[i])
	}
}
