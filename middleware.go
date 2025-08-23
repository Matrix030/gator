package main

import (
	"context"
	"errors"
	"fmt"
	db "github.com/Matrix030/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user db.User) error,
) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return errors.New("not logged in")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to load user %q: %w", s.cfg.CurrentUserName, err)
		}

		return handler(s, cmd, user)
	}
}
