package main

import (
	"context"
	"errors"
	"fmt"

	db "github.com/Matrix030/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user db.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Usage: follow <url>")
	}

	ctx := context.Background()

	URL := cmd.Args[0]
	f, err := s.db.GetFeedByURL(ctx, URL)
	if err != nil {
		return fmt.Errorf("Feed not found for url %q: %w", URL, err)
	}
	params := db.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: f.ID,
	}

	row, err := s.db.CreateFeedFollow(ctx, params)

	if err != nil {
		return fmt.Errorf("Could not follow feed: %w", err)
	}

	fmt.Printf("Followed: %s (by %s)\n", row.FeedName, row.UserName)
	return nil
}
