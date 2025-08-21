package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	ctx := context.Background()

	user, err := s.db.GetUser(ctx, name)
	if err != nil {
		fmt.Println("user not available", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User found and switched successfully!")
	fmt.Printf("debug: %v\n", user)

	return nil
}
