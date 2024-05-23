package minesweeper

import (
	"errors"

	"github.com/jrlmx/repl/internal/repl"
)

func drawCommand(r *repl.Repl, args ...string) error {
	g := r.Config.(*config).game

	if g == nil || g.gamestate != playing {
		return errors.New("no game in progress")
	}

	return nil
}
