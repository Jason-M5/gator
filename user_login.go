package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/jason-M5/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("expects a single argument, the username")
	}

	name := cmd.arguments[0]
	_, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %v", err)
	}

	err = s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("could not set user: %v", err)
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return fmt.Errorf("expects one argument, the username")
	}

	name := cmd.arguments[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("unable to create user in database: %v\n", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %v", err)
	}

	fmt.Printf("User '%s' was successfully created!\n", user.Name)
	fmt.Printf("Logged Data: %+v\n", user)
	return nil

}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset users: %v", err)
	}
	fmt.Println("Users have been reset")
	return nil
}