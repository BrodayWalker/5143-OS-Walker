package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"regexp"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["ls"] = Ls
}

// Ls lists all files and directories in a specified folder. Currently, no
// flags are supported.
func Ls(args []string) {
	argList, flags := ArgSplitter(args)
	var path string = ""

	// ReadDir reads the directory named by dirname and returns a list of
	// directory entries sorted by filename. The entries are a FileInfo
	// object with the following format:
	/*
			type FileInfo interface {
		    	Name() string       // base name of the file
		    	Size() int64        // length in bytes for regular files; system-dependent for others
		    	Mode() FileMode     // file mode bits
		    	ModTime() time.Time // modification time
		    	IsDir() bool        // abbreviation for Mode().IsDir()
				Sys() interface{}   // underlying data source (can return nil)
			}
	*/

	// If no arguments, list all file and folder names only
	if len(argList) == 0 && len(flags) == 0 {
		defaultPrint(".", false)
	} else if len(argList) > 0 && len(flags) == 0 {
		// Different path than current working directory, no flags passed
		path = BuildPathToDir(argList[0])
		defaultPrint(path, false)
	} else {
		// Print files in current folder
		if len(argList) == 0 {
			path = "."
		} else {
			// Different path
			path = argList[0]
		}
		// Build path to directory
		path = BuildPathToDir(path)

		// Find which flags are present (returns a boolean for each flag)
		lFlag, hFlag, aFlag := parseFlags(flags)

		// It's ugly but it handles all the flag permutations
		if lFlag == true && hFlag == false && aFlag == false{
			// -l
			longPrint(path, false, false)
		} else if lFlag == true && hFlag == true && aFlag == false {
			// -lh
			longPrint(path, true, false)
		} else if lFlag == true && hFlag == false && aFlag == true {
			// -la
			longPrint(path, false, true)
		} else if lFlag == true && hFlag == true && aFlag == true {
			// -lha
			longPrint(path, true, true)
		} else if lFlag == false && hFlag == false && aFlag == true {
			// -a
			defaultPrint(path, true)
		} else if lFlag == false && hFlag == true && aFlag == false {
			// -h
			defaultPrint(path, false)
		} else if lFlag == false && hFlag == true && aFlag == true {
			// -ah
			defaultPrint(path, true)
		}
	}
}

// parseFlags checks for -l and -h
func parseFlags(flags []string) (bool, bool, bool) {
	lFlag := false
	hFlag := false
	aFlag := false
	// Loop through flag array
	for _, v := range flags {
		if v == "l" {
			lFlag = true
		}
		if v == "h" {
			hFlag = true
		}
		if v == "a" {
			aFlag = true
		}
	}
	return lFlag, hFlag, aFlag
}

// defaultPrint is the vanilla ice cream of prints. It lists files
// and directories.
func defaultPrint(path string, all bool) {
	hiddenFile := regexp.MustCompile(`^[.]+`)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	// -a prints all files, including hidden files
	if all == true {
		for _, file := range files {
			fmt.Printf(file.Name() + " ")
		}
	} else {
		// no -a, don't print hidden files
		for _, file := range files {
			if !hiddenFile.MatchString(file.Name()) {
				fmt.Printf(file.Name() + " ")
			}
		}
	}
	
	fmt.Printf("\n")
}

// longPrint prints a long listing. If the second argument is true, a long
// listing is printed with sizes in human-readable format.
func longPrint(path string, human bool, all bool) {
	files, _ := ioutil.ReadDir(path)
	hiddenFile := regexp.MustCompile(`^[.]+`)

	// Long listing, not human readable
	if human == false {
		if all == true {
			// Print all files 
			for _, file := range files {
				// Print permissions
				fmt.Printf("%s ", file.Mode())
				// Print owner fields
				// This is not implemented as Windows returns -1 for
				// the group and owner fields
				// Print size
				fmt.Printf("%12d ", file.Size())
				// Print date
				t := file.ModTime()
				fmt.Printf("%v ", t.Format(time.UnixDate))
				// Print file/folder name
				fmt.Printf(file.Name() + " ")
				fmt.Printf("\n")
			}
		} else if all == false {
			// no -a, don't print hidden files
			for _, file := range files {
				if !hiddenFile.MatchString(file.Name()) {
					// Print permissions
				fmt.Printf("%s ", file.Mode())
				// Print owner fields
				// This is not implemented as Windows returns -1 for
				// the group and owner fields
				// Print size
				fmt.Printf("%12d ", file.Size())
				// Print date
				t := file.ModTime()
				fmt.Printf("%v ", t.Format(time.UnixDate))
				// Print file/folder name
				fmt.Printf(file.Name() + " ")
				fmt.Printf("\n")
				}
			}
		}
	} else if human == true {
		if all == true {
			// Long listing, human readable
			for _, file := range files {
				// Print permissions
				fmt.Printf("%s ", file.Mode())

				// Print owner fields
				// This is not implemented as Windows returns -1 for
				// the group and owner fields

				// Convert file size format
				formatSize := divide(file.Size())
				// Print size
				fmt.Printf("%6s ", formatSize)

				// Print date
				t := file.ModTime()
				fmt.Printf("%v ", t.Format(time.UnixDate))
				// Print file/folder name
				fmt.Printf(file.Name() + " ")
				fmt.Printf("\n")
			}
		} else if all == false {
			// no -a, don't print hidden files
			for _, file := range files {
				if !hiddenFile.MatchString(file.Name()) {
					// Print permissions
				fmt.Printf("%s ", file.Mode())
				// Print owner fields
				// This is not implemented as Windows returns -1 for
				// the group and owner fields

				// Convert file size format
				formatSize := divide(file.Size())
				
				// Print size
				fmt.Printf("%6s ", formatSize)
				// Print date
				t := file.ModTime()
				fmt.Printf("%v ", t.Format(time.UnixDate))
				// Print file/folder name
				fmt.Printf(file.Name() + " ")
				fmt.Printf("\n")
				}
			}
		}
	}
}

// divide finds the number of divisions needed to get a file size below
// 1024. Example: if a file size in bytes takes 3 divisions to get below 1024,
// the file is 1 or more gigabytes. The returned string is formatted with some
// order of magnitude (ex: 1GB)
func divide(size int64) string {
	count := 0
	for size > 1024 {
		size /= 1024
		count++
	}

	// Convert size to a string
	fSize := strconv.Itoa(int(size))

	switch count {
	case 0:
		fSize += "B"
	case 1:
		fSize += "KB"
	case 2:
		fSize += "MB"
	case 3:
		fSize += "GB"
	case 4:
		fSize += "TB"
	case 5:
		fSize += "PB"
	}

	return fSize
}
