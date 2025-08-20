package main

import (
	"fmt"
	"github.com/Matrix030/gator/internal/config"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	// 1) Check for at least 2 args: program name + command
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: not enough arguments\nusage: gator <command> [args]")
		os.Exit(1)
	}

	// 2) Read config from file
	cfg, err := config.Read()
	if err != nil {
		// config read failure is a hard error
		log.Fatalf("error reading config: %v", err)
	}

	// 3) Build state
	s := &State{cfg: &cfg}

	// 4) Register commands
	cmds := &Commands{}
	cmds.register("login", handlerLogin)
	// 5) Split CLI input into command name + args
	name := os.Args[1]
	args := os.Args[2:] // everything after the command name

	// 6) Run command, print any error, exit 1 on failure
	if err := cmds.run(s, Command{Name: name, Args: args}); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
