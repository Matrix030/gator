package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsersList(s *state, cmd command) error {
	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return errors.New("couldn't get all users")
	}

	fmt.Println("All users:")
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil

}
