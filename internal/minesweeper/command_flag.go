package minesweeper

import (
	"errors"
	"strconv"

	"github.com/jrlmx/repl/internal/repl"
)

func flagCommand(r *repl.Repl, args ...string) error {
	if len(args) < 2 {
		return errors.New("missing x and/or y")
	}

	x, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("invalid x")
	}

	y, err := strconv.Atoi(args[1])

	if err != nil {
		return errors.New("invalid y")
	}

	g := r.Config.(*config).game

	if g == nil || g.gamestate != playing {
		return errors.New("no game in progress")
	}

	err = g.Flag(x, y)

	if err != nil {
		return err
	}

	return nil
}
