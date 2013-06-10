package main

import (
	"fmt"
	"github.com/GeertJohan/go.linenoise"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to go.linenoise example.")
	writeHelp()
	for {
		str := linenoise.Line("prompt> ")
		fields := strings.Fields(str)

		// check if there is any valid input at all
		if len(fields) == 0 {
			writeUnrecognized()
			continue
		}

		// switch on the command
		switch fields[0] {
		case "help":
			writeHelp()
		case "echo":
			fmt.Printf("echo: %s\n\n", str[5:])
		case "multiline":
			fmt.Println("Setting linenoise to multiline")
			linenoise.SetMultiline(true)
		case "singleline":
			fmt.Println("Setting linenoise to singleline")
			linenoise.SetMultiline(false)
		case "addHistory":
			if len(str) < 12 {
				fmt.Println("No argument given.")
			}
			linenoise.AddHistory(str[11:])
		case "quit":
			fmt.Println("Thanks for running the go.linenoise example.")
			fmt.Println("")
			os.Exit(0)
		default:
			writeUnrecognized()
		}
	}
}

func writeHelp() {
	fmt.Println("help              write this message")
	fmt.Println("echo ...          echo the arguments")
	fmt.Println("multiline         set linenoise to multiline")
	fmt.Println("singleline        set linenoise to singleline")
	fmt.Println("addHistory ...    add arguments to linenoise history")
	fmt.Println("quit              stop the program")
	fmt.Println("")
	fmt.Println("Use the arrow up and down keys to walk through history.")
	fmt.Println("Note that you have to use addHistory to create history entries. Commands are not added to history in this example.")
	fmt.Println("")
}

func writeUnrecognized() {
	fmt.Println("Unrecognized command. Use 'help'.")
}
