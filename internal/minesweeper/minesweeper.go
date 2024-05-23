package minesweeper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jrlmx/repl/internal/repl"
)

type winstate int

const (
	playing winstate = iota
	won
	lost
)

type config struct {
	game *Game
}

func StartCommand(r *repl.Repl, args ...string) error {
	cfg := &config{
		game: nil,
	}

	repl.NewRepl(" minesweeper> ", cfg, getCommands(), true).Start()

	return nil
}

func getCommands() map[string]repl.Command {
	return map[string]repl.Command{
		"new": {
			Name:        "new <size> <mines>",
			Description: "Start a new game",
			Action:      newCommand,
		},
	}
}

func render(g *Game) {
	w := len(strconv.Itoa(g.size * g.size))
	vdiv := strings.Repeat(" ", w) + "|"
	hdiv := "-" + strings.Repeat("-", w) + "-" + strings.Repeat("-", g.size*3)

	for y, row := range g.minefield {
		for x, s := range row {
			if x == 0 {
				fmt.Printf(" %v | ", y)
			}
		}
	}
}
