package tictactoe

import (
	"fmt"

	"github.com/jrlmx/repl/internal/repl"
)

type player struct {
	human bool
	name  string
}

type game struct {
	players       [2]player
	board         [9]string
	currentPlayer *player
}

type config struct {
	game *game
}

func newGame() *game {
	players := [2]player{
		{human: true, name: "X"},
		{human: false, name: "O"},
	}

	return &game{
		players: players,
		board: [9]string{
			" ", " ", " ",
			" ", " ", " ",
			" ", " ", " ",
		},
		currentPlayer: &players[0],
	}
}

func StartCommand(root *repl.Repl, args ...string) error {
	cfg := &config{}

	r := repl.NewRepl(" tictactoe> ", cfg, getCommands(), false)

	err := r.Start()

	if err != nil {
		return err
	}

	return nil
}

func getCommands() map[string]repl.Command {
	return map[string]repl.Command{
		"new": {
			Name:        "new",
			Description: "Start a new game",
			Hooks:       true,
			Action:      newCommand,
		},
	}
}

func render(g *game) {
	div := "---+---+---"

	for i, v := range g.board {
		if v == " " {
			fmt.Printf(" %d ", i+1)
		} else {
			fmt.Printf(" %s ", v)
		}

		if (i+1)%3 == 0 {
			fmt.Println()

			if i < 8 {
				fmt.Println(div)
			}
		} else {
			fmt.Print("|")
		}
	}
}
