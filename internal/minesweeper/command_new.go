package minesweeper

import (
	"github.com/jrlmx/repl/internal/repl"
)

func newCommand(r *repl.Repl, args ...string) error {
	size, mines, err := parseXYInput("size", "mines", args...)

	if err != nil {
		return err
	}

	g, err := NewGame(size, mines)

	if err != nil {
		return err
	}

	r.Config.(*config).game = g

	return nil
}
