package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	u, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	rows, err := s.db.GetFeedFollowsForUser(ctx, u.ID)
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
