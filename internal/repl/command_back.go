package repl

func backCommand(r *Repl, args ...string) error {
	r.stop()

	return nil
}
