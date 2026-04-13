package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Go-Do is a CLI utility for efficient task management.")
		fmt.Println("Usage:\n  godo [command]")
		fmt.Println("Available Commands:")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("  --version  version for godo")
		os.Exit(0)
	}

	args := os.Args[1:]

	if args[0] == "--version" {
		fmt.Println("godo v0.0.0")
	}
}
