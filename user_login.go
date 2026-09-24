package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("expects a single argument, the username")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("could not set user: %v", err)
	}
	fmt.Println("User has been set")
	return nil
}
