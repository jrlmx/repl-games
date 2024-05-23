package minesweeper

import (
	"errors"

	"github.com/jrlmx/repl/internal/repl"
)

func flagCommand(r *repl.Repl, args ...string) error {
	x, y, err := parseXYInput("x", "y", args...)

	if err != nil {
		return err
	}

	cfg := r.Config.(*config)

	if cfg.gameInProgress() {
		return errors.New("no game in progress")
	}

	err = cfg.game.Flag(x, y)

	if err != nil {
		return err
	}

	return nil
}
