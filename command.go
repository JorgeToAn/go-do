package main

import (
	"errors"
	"fmt"
)

type command struct {
	Name        string
	Args        []string
	Description string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.handlers[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return errors.New("command not found")
	}

	return handler(s, cmd)
}

func printCommands() {
	cmds := availableCommands()
	for _, cmd := range cmds {
		fmt.Printf("  %s  %s\n", cmd.Name, cmd.Description)
	}
}

func availableCommands() []command {
	cmds := []command{
		{
			Name:        "configure",
			Description: "sets the config file values",
		},
		{
			Name:        "register",
			Description: "creates a new user",
		},
		{
			Name:        "login",
			Description: "change the current user",
		},
		{
			Name:        "help",
			Description: "prints this help message",
		},
		{
			Name:        "version",
			Description: "prints the go-do cli version",
		},
	}
	return cmds
}
