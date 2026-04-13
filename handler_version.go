package main

import "fmt"

func handlerVersion(s *state, cmd command) error {
	fmt.Printf("go-do %s\n", s.version)
	return nil
}
