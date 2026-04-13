package main

import "fmt"

func handlerHelp(s *state, cmd command) error {
	printHelp()
	return nil
}

func printHelp() {
	fmt.Println("Go-Do is a CLI utility for efficient task management.")
	fmt.Println("Usage:\n  go-do [command]")
	fmt.Println("Available Commands:")
	printCommands()
}
