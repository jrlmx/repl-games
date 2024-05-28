package tictactoe

import "github.com/jrlmx/repl/internal/repl"

func newCommand(r *repl.Repl, args ...string) error {
	cfg := r.Config.(*config)

	cfg.game = newGame()

	return nil
}
