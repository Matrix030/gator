package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		fmt.Println("An error occured while deleting users:", err)
		os.Exit(1)
	}
	fmt.Println("All users deleted.")
	return nil
}
