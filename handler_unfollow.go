package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	db "github.com/Matrix030/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user db.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: unfollow <URL>")
	}

	url := cmd.Args[0]

	ctx := context.Background()
	_, err := s.db.UnfollowFeedByUserAndUrl(ctx, db.UnfollowFeedByUserAndUrlParams{
		UserID: user.ID,
		Url:    url,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("You are not following %s", url)
		}
		return fmt.Errorf("Failed to unfollow %s: %w", url, err)
	}

	fmt.Printf("Unfollowed %s\n", url)
	return nil
}
