package minesweeper

import (
	"errors"
	"strconv"

	"github.com/jrlmx/repl/internal/repl"
)

func newCommand(r *repl.Repl, args ...string) error {
	if len(args) < 2 {
		return errors.New("missing size and/or mines")
	}

	size, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("invalid size")
	}

	mines, err := strconv.Atoi(args[1])

	if err != nil {
		return errors.New("invalid mines")
	}

	g, err := NewGame(size, mines)

	if err != nil {
		return err
	}

	r.Config.(*config).game = g

	return nil
}
