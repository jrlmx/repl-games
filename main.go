package main

import (
	"github.com/jrlmx/repl/internal/minesweeper"
	"github.com/jrlmx/repl/internal/repl"
)

func main() {
	r := repl.NewRepl(" main menu> ", nil, getCommands(), true)
	err := r.Start()

	if err != nil {
		panic(err)
	}
}

func getCommands() map[string]repl.Command {
	return map[string]repl.Command{
		"minesweeper": {
			Name:        "minesweeper",
			Description: "Start a new minesweeper game",
			Action:      minesweeper.StartCommand,
		},
	}
}
