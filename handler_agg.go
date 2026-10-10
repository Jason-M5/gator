package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	fmt.Println("Aggregating feeds...")
	feeds, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Println(feeds)
	return nil
}
