package repl

import (
	"os"
)

func exitCommand(r *Repl, args ...string) error {
	os.Exit(0)

	return nil
}
