package repl

import "fmt"

func helpCommand(r *Repl, args ...string) error {
	for _, cmd := range r.commands {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Description)
	}

	return nil
}
