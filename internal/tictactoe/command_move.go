package tictactoe

import (
	"errors"

	"github.com/jrlmx/repl/internal/repl"
	"github.com/jrlmx/repl/internal/utils"
)

func moveCommand(r *repl.Repl, args ...string) error {
	cfg := r.Config.(*config)

	if cfg.game == nil {
		return errors.New("no game in progress")
	}

	x, err := utils.ParseXInput("x", args...)

	if err != nil {
		return err
	}

	err = cfg.game.move(x)

	if err != nil {
		return err
	}

	return nil
}
