package config

import (
	"errors"
	"fmt"
)

type State struct {
	configPointer *Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	funcMap map[string]func(*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	handler, ok := c.funcMap["run"]

}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Usage: gator login <username>")
	}
	userName := cmd.Args[0]

	err := s.configPointer.SetUser(cmd.Args[1])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %q\n", userName)
	return nil

}
