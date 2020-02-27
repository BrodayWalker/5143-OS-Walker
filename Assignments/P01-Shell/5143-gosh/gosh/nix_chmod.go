// +build !windows

package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strconv"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["chmod"] = Chmod
}

func Chmod(args []string) {

	strmode := args[0]
	filename := args[1]
	mode, err := strconv.Atoi(strmode)
	if err != nil {
		fmt.Println("Error with mode: ", err)
	}

	err = unix.Chmod(filename, uint32(mode))
	if err != nil {
		fmt.Println("Error changing permissions: ", err)
	}

}
