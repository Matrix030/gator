package main

import (
	"database/sql"
	"github.com/Matrix030/gator/internal/config"
	"github.com/Matrix030/gator/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//database stuff
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	//commands to be regitered
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//registering commands
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegisterUsers)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsersList)
	cmds.register("agg", handlerAggregation)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeedsList)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
