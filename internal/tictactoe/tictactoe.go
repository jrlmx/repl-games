package tictactoe

import "github.com/jrlmx/repl/internal/repl"

type config struct {
	//
}

func StartCommand() {
	r := repl.NewRepl(" tictactoe> ", &config{}, getCommands(), false)

	r.Start()
}

func getCommands() map[string]repl.Command {
	return map[string]repl.Command{
		// "new": {
		// 	Name:        "new",
		// 	Description: "Start a new game",
		// 	Hooks:       true,
		// 	Action:      newCommand,
		// },
		// "draw": {
		// 	Name:        "draw",
		// 	Description: "Draw the current game",
		// 	Hooks:       true,
		// 	Action:      drawCommand,
		// },
		// "move": {
		// 	Name:        "move <x> <y>",
		// 	Description: "Make a move",
		// 	Hooks:       true,
		// 	Action:      moveCommand,
		// },
	}
}
