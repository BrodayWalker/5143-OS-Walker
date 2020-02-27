package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// will be used to count number of instances of the first character of each line
var numDict = map[string]int{"0": 0, "1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0}
var lowerDict = map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
var upperDict = map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "E": 0, "F": 0, "G": 0, "H": 0, "I": 0, "J": 0, "K": 0, "L": 0, "M": 0, "N": 0, "O": 0, "P": 0, "Q": 0, "R": 0, "S": 0, "T": 0, "U": 0, "V": 0, "W": 0, "X": 0, "Y": 0, "Z": 0}

func init() {
	// Add this command's function to the command mapping
	ComMap["sort"] = Sort
}

// Sort performs a brute-force algorithm to sort each line of an input file
// by the first character. It currently does not sort past the first character
// Usage:
//				sort input1 input2 inputN
func Sort(args []string) {
	argsSize := len(args)
	if argsSize == 1 {
		sortFile(args[0]) // call sorting function with input file; standard procedure
	} else if argsSize > 1 {
		args = append(args, ">") // append the redirection arrow and a temporary file to args
		args = append(args, "sort.tmp")
		Cat(args)                  // have Cat concatenate all the files and stick them into the sort.tmp file
		sortFile(args[argsSize-1]) // call sorting function with temporary file, sort.tmp
	} else {
		fmt.Println("Error: no input arguments. Use form 'sort input1 input2 inputN'")
	}
}

// Since sort on bash doesn't use the ascii order, regular sorting doesn't work.
// This function will search the input files four times, independently searching
// for either numbers, lowercase letters, uppercase letters, or symbols.
// It then sorts each case and dumps them into Stdout in the order symbols, numbers,
// lowercase letters, uppercase letters.
func sortFile(file string) {
	var (
		numbers, lowers, uppers, symbols = []string{}, []string{}, []string{}, []string{}
		element                          int
	)
	inputFile, _ := os.OpenFile(file, os.O_RDONLY, 0755)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// we check if Scan() read in a blank line. If so, skip it
		if scanner.Text() != "" {
			// If the first character of the line is a number, add it to the number array
			if _, numIn := numDict[string(scanner.Text()[0])]; numIn {
				numbers = append(numbers, scanner.Text())
				// If the first character of the line is a lowercase letter, add it to the number array
			} else if _, lowerIn := lowerDict[string(scanner.Text()[0])]; lowerIn {
				lowers = append(lowers, scanner.Text())
				// If the first character of the line is a uppercase letter, add it to the number array
			} else if _, upperIn := upperDict[string(scanner.Text()[0])]; upperIn {
				uppers = append(uppers, scanner.Text())
				// everything else, add it to the number array
			} else {
				symbols = append(symbols, scanner.Text())
			}
		}
	}
	os.Remove("sort.tmp")
	inputFile.Close()
	// sort each of the four arrays
	sort.Strings(symbols)
	sort.Strings(numbers)
	sort.Strings(lowers)
	sort.Strings(uppers)
	// print them all to Stdout in the bash way of sorting:
	// symbols, numbers, lowers, then uppers
	for element = range symbols {
		fmt.Println(symbols[element])
	}
	for element = range numbers {
		fmt.Println(numbers[element])
	}
	for element = range lowers {
		fmt.Println(lowers[element])
	}
	for element = range uppers {
		fmt.Println(uppers[element])
	}
}
