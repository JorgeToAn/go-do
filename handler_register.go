package main

import (
	"context"
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: go-do %s username", cmd.Name)
	}

	username := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't create user: %s", err)
	}

	fmt.Println("User created:")
	fmt.Printf("  ID:   %s\n", user.ID)
	fmt.Printf("  Name: %s\n", user.Name)
	return nil
}
