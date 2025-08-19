package main

import (
	"errors"
	"fmt"
	"github.com/Matrix030/gator/internal/config"
)

type State struct {
	cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	funcMap map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	if c.funcMap == nil {
		c.funcMap = make(map[string]func(*State, Command) error)
	}
	c.funcMap[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	h, ok := c.funcMap[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return h(s, cmd)
}

func handlerLogin(s *State, cmd Command) error {
	if s == nil || s.cfg == nil {
		return errors.New("internal error: state/config not initialized")
	}
	if len(cmd.Args) != 1 {
		return errors.New("usage: gator login <username>")
	}
	username := cmd.Args[0]

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("User has been set to %q\n", username)
	return nil
}
