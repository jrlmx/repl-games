package tictactoe

import (
	"errors"

	"github.com/jrlmx/repl/internal/repl"
)

func drawCommand(r *repl.Repl, args ...string) error {
	cfg := r.Config.(*config)

	if cfg.game == nil {
		return errors.New("no game in progress")
	}

	// Here we don't need to do anything, just return nil
	// The Render After hook will take care of rerendering the game

	return nil
}
