package repl

func clearCommand(r *Repl, args ...string) error {
	Clear()

	return nil
}
