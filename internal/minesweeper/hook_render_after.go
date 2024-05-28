package minesweeper

import (
	"fmt"

	"github.com/jrlmx/repl/internal/repl"
)

func renderAfterHook(r *repl.Repl, args ...string) error {
	cfg := r.Config.(*config)

	if cfg.game == nil {
		return nil
	}

	repl.Clear()

	render(cfg.game)

	if cfg.game.gamestate == won {
		fmt.Println("You win! Start a new game to continue...")
	}

	if cfg.game.gamestate == lost {
		fmt.Println("You lose! Start a new game to continue...")
	}

	return nil
}
