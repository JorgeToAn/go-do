package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/JorgeToAn/go-do/internal/config"
)

//go:embed version.txt
var cliVersion string

type state struct {
	config  *config.Config
	version string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := state{
		config:  &cfg,
		version: cliVersion,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("help", handlerHelp)
	cmds.register("version", handlerVersion)

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.run(&s, cmd); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
