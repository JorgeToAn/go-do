package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/JorgeToAn/go-do/internal/config"
	"github.com/JorgeToAn/go-do/internal/database"
	_ "github.com/lib/pq"
)

//go:embed version.txt
var cliVersion string

type state struct {
	config  *config.Config
	version string
	db      *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("couldn't connect to database: %s", err)
	}
	dbQueries := database.New(db)

	s := state{
		config:  &cfg,
		version: cliVersion,
		db:      dbQueries,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("configure", handlerConfigure)
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
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
