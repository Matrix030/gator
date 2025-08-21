package main

import (
	"context"
	"fmt"
	db "github.com/Matrix030/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handlerRegisterUsers(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	//create a new user in the DB
	ctx := context.Background()
	now := time.Now().UTC()

	params := db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	}

	_, err := s.db.CreateUser(ctx, params)
	if err != nil {
		fmt.Println("could not create user:", err)
		os.Exit(1)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return err
	}

	fmt.Printf("User %q created\n", name)
	return nil
}
