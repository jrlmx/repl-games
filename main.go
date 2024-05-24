package main

import (
	"log"

	"github.com/jrlmx/repl/internal/minesweeper"
	"github.com/jrlmx/repl/internal/repl"
	"github.com/jrlmx/repl/internal/tictactoe"
)

func main() {
	r := repl.NewRepl(" main menu> ", nil, getCommands(), true)
	err := r.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func getCommands() map[string]repl.Command {
	return map[string]repl.Command{
		"minesweeper": {
			Name:        "minesweeper",
			Description: "Start a new minesweeper game",
			Action:      minesweeper.StartCommand,
		},
		"tictactoe": {
			Name:        "tictactoe",
			Description: "Start a new tictactoe game",
			Action:      tictactoe.StartCommand,
		},
	}
}
