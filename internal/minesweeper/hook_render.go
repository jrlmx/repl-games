package minesweeper

import "github.com/jrlmx/repl/internal/repl"

func renderHook(r *repl.Repl, args ...string) error {
	if r.Config.(*config).game == nil {
		return nil
	}

	render(r.Config.(*config).game)

	return nil
}
