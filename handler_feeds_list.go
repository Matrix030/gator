package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeedsList(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return errors.New("Usage: feeds")
	}
	ctx := context.Background()

	allFeedList, err := s.db.ListFeedsWithUsers(ctx)
	if err != nil {
		return errors.New("Couldn't get all feeds")
	}

	fmt.Println("All Feeds:")
	for _, feeds := range allFeedList {
		fmt.Printf("Feed Name: %v\n", feeds.FeedName)
		fmt.Printf("URL: %v\n", feeds.Url)
		fmt.Printf("Creator Name: %v\n", feeds.UserName)
	}

	return nil
}
