package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// The helpDict hold all commands' help info to be printed whenever 'cmd --help' is entered in the shell
//--------------------------------------------------------------------------------------*-----------------------------------------------------------------------------------------*-----------------------------------------------------------------------------------------*-----------------------------------------------------------------------------------------*-----------------------------------------------------------------------------------------*
var helpDict = map[string]string{
	"cat":         "Cat will either concatenate a file and print it to std out or intake\ntwo or more files and print them as if they were concatenated. If Cat was called without\nany arguments, it will take input from Stdin and just output it to Stdout.\nUsage:\n\tcat\n\tcat file\n\tcat file1 file2 fileN",
	"cd":          "Cd changes directories.\nUsage:\n\tcd path/to/directory\n\tcd ..\n\tcd ~",
	"cp":          "Cp will copy a file or a set of files into another location. This function only\nworks if the folders on the destination path exist and the paths are Unix-formatted,\ni.e.\n\tpath/to/destination\n\t\tNOT\n\tpath\\to\\destination OR path\\\\to\\\\destination\nUsage:\n\tcp path/to/file path/to/copy/to\n\tcp path/to/file1 path/to/file2 path/to/fileN path/to/copy/to",
	"chmod":       "Chmod changes the mode/permissions of a file or directory.\nUsage:\n\tchmod file mode",
	"echo":        "Echo repeats what is typed after the command.\nUsage:\n\techo statement",
	"exclamation": "Exclamation will return the n-th command from the commandList and execute\nit automatically. If the command does not exist, it returns an error message and the\nshell continues like normal.\nIn this function, location is the command number that follows the\n'!' character.\nUsage:\n\t![instruction #]",
	"gosh":        "Gosh is a sophisticated shell simulation developed in GoLang by\n\tBroday Walker (team leader)\n\tJeremy Glebe\n\tCorbin Matamoros\nfor the Advanced Operating System class at Midwestern State University.",
	"grep":        "",
	"head":        "Head prints the first 10 lines of a file or the number of lines\nspecified after the -n flag.\nUsage:\n\thead file\n\thead -n [# of lines] file",
	"help":        "Help prints command and command usage information.\nUsage:\n\thelp command",
	"history":     "History prints all commands executed (except Exit), syntactically correct or not.\nUsage:\n\thistory",
	"less":        "Less will print out the contents of an input file one page at a time,\nwrapping around words and in between words if the input file line has more characters\nthan the terminal has width.\nUsage:\n\tless <inputfile>",
	"ls":          "Ls lists all files and directories in the currect folder.\nUsage:\n\tls\n\tls -a\n\tls -l\n\tls -h",
	"mkdir":       "Mkdir makes a new directory. Currently, the only supported functionality\nincludes making a folder in the current directory with a folder name supplied by the\nuser as the first argument.\nUsage\n\tmkdir path/to/directory/directory_name",
	"mv":          "Mv moves the file given in the first argument to the location of the\nsecond argument. Mv also renames files.\nUsage:\n\tmv path/to/file path/to/desired/location",
	"nix_mkdir":   "",
	"mix_mv":      "",
	"pipe":        "Pipeline takes a list of Commands, then for each Command, sets its output\nto a file and uses that file as the first arg of the next command, executing them all\nin sequence.\n\nNotes:\n\t- Actual command calling is handled in Execute()\n\t- In the event of an input redirection in use with Pipes, redirect is\n\thandled by Execute(), which will in turn call RedirectAndExecute()\nUsage:\n\thistory.go | wc\n\tcat file1 file2 > outfile | head -n 10 outfile",
	"pwd":         "Pwd prints the working directory.\nUsage:\n\tpwd",
	"rm":          "Rm removes a folder or file.\nUsage:\n\trm file\n\trm -r folder\n\trm -rf folder",
	"sort":        "Sort takes a file or set of files and sorts them by line.\nUsage:\n\rsort file1\n\tsort file1 file2 fileN",
	"split":       "Split will divide an input file into a certain number of individual files\ndepending on arguments. If no arguments are sent in the file is divided into 10 parts.\nIf -l <n> is used, the file is divided into parts with n number of lines each.\nUsage:\n\tsplit [option](-l) [option value](<n>) [input](infile)",
	"tail":        "Tail prints the final n number of lines of a file. By defualt, n is set\nto 10. At some point, build a way to handle being passed an absolute path (split on '\\').\nCurrently only works if the file is in the current working directory.\nUsage:\n\ttail file\n\ttail -n [# of lines] file",
	"touch":       "Touch changes the access and/or modification timestamps of the specified files.\nUsage:\n\ttouch file",
	"wc":          "Wc will output the number of characters, words, and lines in a file.\nUsage:\n\twc file"}

var (
	// list that will hold all commands typed in the terminal
	commandList []string
	// tmpfile will hold a command queried by the "!" character.
	// oldStdin will remember the traditional Stdin when it is
	// switched to tmpfile (see 'exclamation.go').
	// Checker is a flag that will "activate" code in 'getInput()'
	// to run if "!" is used.
	tmpfile  *os.File
	oldStdin *os.File = os.Stdin
	checker  bool     = false
	// saves the position in commandHistory list where this session's
	// commands start. That way, when the shell closes, we only copy this
	// session's commands to the history file.
	whereLeftOff int
	// ComMap with be used to create a map of strings (command keys) to Commands (the functions)
	ComMap = make(map[string]CommandFunc)
)

type (
	// CommandFunc will be used to treat functions of the given format as a type "Command"
	CommandFunc func(args []string)
	// Command : A command grouped with arguments for calling it
	Command struct {
		key        string
		args       []string
		infile     *string
		outfile    *string
		appendMode bool
	}
)

func main() {
	// loads the command history file contents from gosh_history.tmp into RAM
	loadHistory()
	// Main loop
	for {
		// Print prompt
		// The double percent sign must be used to print a literal percent sign
		fmt.Printf("%% ")
		// Get the user's input
		input := getInput()
		//if carriage return is hit without anything typed in, restart the loop
		if input == "" {
			continue
		}
		// loads a command into commandList array
		commandList = append(commandList, input)
		// Split the input by instances of && (multiple commands to run)
		commandLines := strings.Split(input, "&&")
		// Run each command line
		for _, line := range commandLines {
			// Split each line of input by pipes if necessary
			piping := strings.Split(line, "|")
			// If there is any piping, we'll need to make a Command list
			if len(piping) > 1 {
				var pipe []Command
				for _, command := range piping {
					pipe = append(pipe, parseCommand(command))
				}
				PipeLine(pipe)
			} else {
				// If there is no piping, we just need to run the command
				// Get the command from the line of text
				command := parseCommand(line)
				// Standardize command
				command.key = strings.ToLower(command.key)
				// Execute the command in the standard way
				Execute(command)
			}
		}
	}
}

// loads the contents of gosh_history.tmp into the commandList array
func loadHistory() {
	// open the gosh_history file for reading
	historyFile, _ := os.OpenFile("gosh_history.tmp", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0755)
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		// append the command from the gosh_history.tmp and remove the
		// \n
		commandList = append(commandList, strings.TrimRight(scanner.Text(), "\n"))
	}
	historyFile.Close()
	// remember where the commands for this current session begin
	whereLeftOff = len(commandList)
}
func getInput() string {
	var (
		// line will hold the command
		// e will hold any error messages for us to ignore
		line string
		e    error
	)
	// If checker is true, it means the exclamation command was run, so we need
	// to run the command it saved to its temporary file. Therefore,
	// temporarily change Stdin to the temporary file created in exclamation.go,
	// read in its contents, and execute the result
	if checker {
		checker = false
		//switch to the new Stdin - the temporary file
		os.Stdin = tmpfile
	}
	// Create a keyboard reader, which will either read from the keyboard
	// or from the temporary file created in the exclamation.go
	keyboard := bufio.NewReader(os.Stdin)
	// Read a line of input from Stdin until carriage return
	line, e = keyboard.ReadString('\n')
	// if Stdin was changed, switch back to the original Stdin
	// if it wasn't changed, oh well. Doing it this way makes for a cleaner
	//code
	os.Stdin = oldStdin
	// Print out any errors
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
	// Trim \r\n for Windows
	if runtime.GOOS == "windows" {
		line = strings.TrimRight(line, "\r\n")
	} else {
		line = strings.TrimRight(line, "\n")
	}
	return line
}

func parseCommand(line string) Command {
	// Trim any leading and trailing spaces resulting from '&&' or '|' splits
	// This has to be done to process multiple commands. It just does.
	line = strings.TrimLeft(line, " ")
	line = strings.TrimRight(line, " ")
	// Separate the arguments
	symbols := strings.Split(line, " ")
	command := symbols[0]
	args := symbols[1:]
	// Set all redirection items to nil by default
	var ifile *string = nil
	var ofile *string = nil
	appendMode := false
	// Get any input redirection
	for i, arg := range args {
		if arg == "<" {
			ifile = new(string)
			*ifile = args[i+1]
			// Remove element at i
			copy(args[i:], args[i+2:]) // Shift left 2 indices
			args = args[:len(args)-2]  // Truncate
		}
		if arg == ">" {
			ofile = new(string)
			*ofile = args[i+1]
			// Remove element at i
			copy(args[i:], args[i+2:]) // Shift left 2 indices
			args = args[:len(args)-2]  // Truncate
		} else if arg == ">>" {
			ofile = new(string)
			*ofile = args[i+1]
			appendMode = true
			// Remove element at i
			copy(args[i:], args[i+2:]) // Shift left 2 indices
			args = args[:len(args)-2]  // Truncate
		}
	}
	// Return command and arguments
	return Command{command, args, ifile, ofile, appendMode}
}

// Execute ..
// Runs a command if it is a valid command defined in the ComMap
// or if it is one of a few special commands such as `exit`.
// Notes:
// - In the event of redirection, this function calls a special
//     execution function, RedirectAndExecute()
// - This is not to be confused with piping, which is handled by
//     a completely different function PipeLine() which in turn
//     calls Execute() on each element of the pipeline (after
//     which individual elements may also be redirected)
func Execute(command Command) {
	// If any redirection has occured
	if command.infile != nil || command.outfile != nil {
		// Run with redirected input/output
		RedirectAndExecute(command)
	} else {
		// No redirection
		// Route the command to call the proper function
		if len(command.args) > 0 && command.args[0] == "--help" { // display the commands help info
			help(command.key)
		} else if com, valid := ComMap[command.key]; valid {
			com(command.args)
		} else if string(command.key[0]) == "!" {
			Exclamation(command.key[1:])
		} else if command.key == "exit" {
			saveHistory()
			os.Exit(0)
		} else {
			fmt.Println("Command not found.")
		}
	}
}

// when the shell is about to be closed, saved all this session's commands to the
// gosh_history file.
func saveHistory() {
	// limit is one minus the length of commandList so to not include
	// any exit commands
	limit := len(commandList) - 1
	historyFile, _ := os.OpenFile("gosh_history.tmp", os.O_WRONLY|os.O_APPEND, 0)
	// add only the commands executed during this session to the
	// gosh history file by starting from the point we saved in
	// whereLeftOff
	for i := whereLeftOff; i < limit; i++ {
		historyFile.WriteString(commandList[i] + "\n")
	}
	historyFile.Close()
}

// help will display a command's help info or an error if a unsupported command
// was entered
func help(args string) {
	if info, exists := helpDict[args]; exists {
		fmt.Println(info)
	} else {
		fmt.Println("This command is not supported in this shell.")
	}
}
