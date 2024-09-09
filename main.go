package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  [command]")
}

func printVersion() {
	fmt.Println("Version: 1.0.0")
}
