package main

import (
	"fmt"
	"github.com/GeertJohan/go.linenoise"
	"os"
)

func main() {
	fmt.Println("Welcome to go.linenoise example.")
	writeHelp()
	for {
		str := linenoise.Line("prompt> ")
		switch str {
		case "quit":
			fmt.Println("Thanks for running the go.linenoise example.")
			fmt.Println("")
			os.Exit(0)
		case "help":
			writeHelp()
		default:
			fmt.Printf("echo: %s\n\n", str)
		}
	}
}

func writeHelp() {
	fmt.Println("help: write this message.")
	fmt.Println("quit: stop the program.")
	fmt.Println("")
}
