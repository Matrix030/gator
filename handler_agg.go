package main

import (
	"context"
	"fmt"
)

func handlerAggregation(s *state, cmd command) error {
	var feedURL = "https://www.wagslane.dev/index.xml"

	resFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("RSS struct:\n %+v", resFeed)
	return nil
}
