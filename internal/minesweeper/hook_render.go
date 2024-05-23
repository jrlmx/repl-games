package minesweeper

import (
	"fmt"

	"github.com/jrlmx/repl/internal/repl"
)

func renderAfterHook(r *repl.Repl, args ...string) error {
	if r.Config.(*config).game == nil {
		return nil
	}

	g := r.Config.(*config).game

	render(g)

	if g.gamestate == won {
		fmt.Println("You win! Start a new game to continue...")
	}

	if g.gamestate == lost {
		fmt.Println("You lose! Start a new game to continue...")
	}

	return nil
}
