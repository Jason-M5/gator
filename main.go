package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jason-M5/gator/internal/config"
	"github.com/jason-M5/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("unable to open db: %v", err)
	}

	dbQueries := database.New(db)

	progState := &state{
		db:     dbQueries,
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
