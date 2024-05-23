package minesweeper

import (
	"errors"

	"github.com/jrlmx/repl/internal/repl"
)

func explodeCommand(r *repl.Repl, args ...string) error {
	cfg := r.Config.(*config)

	if !cfg.gameInProgress() {
		return errors.New("no game in progress")
	}

	cfg.game.explode()

	return nil
}
