package main

import (
	"log"
	"os"

	"github.com/jason-M5/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	progState := &state{
		config: &cfg,
	}
	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments")
	}
	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}
	err = cmds.run(progState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
