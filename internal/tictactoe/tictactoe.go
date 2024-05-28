package tictactoe

import (
	"fmt"
	"math/rand"

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
	coinToss := rand.Intn(2)

	var player1 player
	var player2 player

	if coinToss == 0 {
		player1 = player{human: true, name: "X"}
		player2 = player{human: false, name: "O"}
	} else {
		player1 = player{human: false, name: "X"}
		player2 = player{human: true, name: "O"}
	}

	players := [2]player{
		player1,
		player2,
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

func (g *game) move(x int) error {
	if g.board[x] != " " {
		return fmt.Errorf("invalid move: cell %d is already taken", x)
	}

	g.board[x] = g.currentPlayer.name

	return nil
}

func StartCommand(root *repl.Repl, args ...string) error {
	cfg := &config{}

	r := repl.NewRepl(" tictactoe> ", cfg, getCommands(), false)

	r.SetRunAfterAction(renderAfterHook)

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
		"draw": {
			Name:        "draw",
			Description: "Draw the board",
			Hooks:       true,
			Action:      drawCommand,
		},
		"move": {
			Name:        "move <x> <y>",
			Description: "Make a move on the board",
			Hooks:       true,
			Action:      moveCommand,
		},
	}
}

func render(g *game) {
	div := "---+---+---"

	for i, v := range g.board {
		if v == " " {
			fmt.Printf(" %d ", i)
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
