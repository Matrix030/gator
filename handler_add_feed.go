package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	db "github.com/Matrix030/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	if len(cmd.Args) != 2 {
		return errors.New("Usage: addfeed [feed_name] [url]")
	}
	feedName := cmd.Args[0]
	URL := cmd.Args[1]

	params := db.CreateFeedParams{
		Name:   feedName,
		Url:    URL,
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, params)
	if err != nil {
		fmt.Println("Could not create feed: ", err)
		os.Exit(1)
	}

	fmt.Printf("feed %q created\n", feed.ID)
	return nil
}
