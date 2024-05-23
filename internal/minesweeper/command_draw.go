package minesweeper

import (
	"errors"

	"github.com/jrlmx/repl/internal/repl"
)

func drawCommand(r *repl.Repl, args ...string) error {
	if r.Config.(*config).game == nil {
		return errors.New("no game in progress")
	}

	g := r.Config.(*config).game

	if g == nil || g.gamestate != playing {
		return errors.New("no game in progress")
	}

	return nil
}
