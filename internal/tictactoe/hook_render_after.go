package tictactoe

import (
	"github.com/jrlmx/repl/internal/repl"
)

func renderAfterHook(r *repl.Repl, args ...string) error {
	cfg := r.Config.(*config)

	if cfg.game == nil {
		return nil
	}

	repl.Clear()

	render(cfg.game)

	return nil
}
