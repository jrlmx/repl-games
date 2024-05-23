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

func (cfg *config) gameInProgress() bool {
	return cfg.game != nil && cfg.game.gamestate == playing
}

func StartCommand(r *repl.Repl, args ...string) error {
	cfg := &config{
		game: nil,
	}

	sr := repl.NewRepl(" minesweeper> ", cfg, getCommands(), false)

	sr.SetRunAfterAction(renderAfterHook)

	sr.Start()

	return nil
}

func getCommands() map[string]repl.Command {
	return map[string]repl.Command{
		"new": {
			Name:        "new <size> <mines>",
			Description: "Start a new game",
			Hooks:       true,
			Action:      newCommand,
		},
		"hit": {
			Name:        "hit <x> <y>",
			Description: "Hit a square",
			Hooks:       true,
			Action:      hitCommand,
		},
		"flag": {
			Name:        "flag <x> <y>",
			Description: "Flag a square",
			Hooks:       true,
			Action:      flagCommand,
		},
		"draw": {
			Name:        "draw",
			Description: "Draw the board",
			Hooks:       true,
			Action:      drawCommand,
		},
		"explode": {
			Name:        "explode",
			Description: "Explode the board",
			Hooks:       true,
			Action:      explodeCommand,
		},
	}
}

func render(g *Game) {
	maxNumWidth := len(strconv.Itoa(g.size))
	cols := " " + strings.Repeat(" ", maxNumWidth) + "|"
	hdiv := strings.Repeat("-", maxNumWidth+1) + "+" + strings.Repeat("-", g.size*3)

	for i := 0; i < g.size; i++ {
		numStr := strconv.Itoa(i)
		cols += strings.Repeat(" ", maxNumWidth-len(numStr)) + numStr + " "
	}

	board := cols + "\n" + hdiv + "\n"

	for y, row := range g.minefield {
		numWidth := len(strconv.Itoa(y))
		ln := fmt.Sprintf(strings.Repeat(" ", maxNumWidth-numWidth)+"%d |", y)

		for _, s := range row {
			if s.IsTriggered() {
				if s.IsMined() {
					ln += " * "
				} else {
					ln += fmt.Sprintf(" %d ", s.getValue())
				}
			} else if s.IsFlagged() {
				ln += " F "
			} else {
				ln += " _ "
			}
		}

		board += ln + "\n"
	}

	fmt.Println(board)
}
