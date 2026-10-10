package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jason-M5/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	args := len(cmd.Args)
	if args != 2 {
		return fmt.Errorf("Adding a feed takes 2 args, got: %d", args)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("Adding feed", feed)
	return nil
}
