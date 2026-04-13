package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: go-do %s username", cmd.Name)
	}

	username := cmd.Args[0]
	user, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user")
	}

	return s.config.SetUser(user.Name)
}
