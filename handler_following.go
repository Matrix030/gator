package main

import (
	"context"
	"fmt"
	db "github.com/Matrix030/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user db.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	rows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("Could not fetch follows: %w", err)
	}

	if len(rows) == 0 {
		fmt.Printf("You are not following any feeds.")
		return nil
	}

	for _, r := range rows {
		fmt.Println(r.FeedName)
	}
	return nil
}
