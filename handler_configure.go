package main

import "fmt"

func handlerConfigure(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: go-do %s db_url", cmd.Name)
	}

	dbUrl := cmd.Args[0]
	return s.config.SetDbUrl(dbUrl)
}
