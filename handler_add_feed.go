package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	db "github.com/Matrix030/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return errors.New("Usage: addfeed [feed_name] [url]")
	}

	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	feedName := cmd.Args[0]
	URL := cmd.Args[1]

	feed, err := s.db.CreateFeed(ctx, db.CreateFeedParams{
		Name:   feedName,
		Url:    URL,
		UserID: user.ID,
	})
	if err != nil {
		fmt.Println("Could not create feed:", err)
		os.Exit(1)
	}

	_, err = s.db.CreateFeedFollow(ctx, db.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		fmt.Println("Could not create follow:", err)
		os.Exit(1)
	}

	fmt.Printf("Feed %q (%s) created and followed by %s\n", feed.Name, feed.Url, user.Name)
	return nil
}

