package main

import (
	"errors"
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
