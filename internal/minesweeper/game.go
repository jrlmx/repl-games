package minesweeper

import (
	"errors"
	"math/rand"
)

type Game struct {
	minefield              [][]square
	gamestate              winstate
	size, mines, movesLeft int
}

func NewGame(size, mines int) (*Game, error) {
	if size < 3 || size > 30 {
		return nil, errors.New("size must be between 3 and 30")
	}

	if mines < 1 || mines > size*size {
		return nil, errors.New("mines must be between 1 and <size> squared")
	}

	minefield := make([][]square, size)

	for y := 0; y < size; y++ {
		row := make([]square, size)

		for x := 0; x < size; x++ {
			row[x] = newSquare(x, y)
		}

		minefield[y] = row
	}

	m := mines

	for m > 0 {
		mx := rand.Intn(size)
		my := rand.Intn(size)

		s := &minefield[my][mx]

		if s.mined {
			continue
		}

		s.mined = true
		m--

		for _, n := range *s.getNeighbors(size) {
			minefield[n.y][n.x].value++
		}
	}

	return &Game{
		minefield: minefield,
		gamestate: playing,
		size:      size,
		mines:     mines,
		movesLeft: size*size - mines,
	}, nil
}

func (g *Game) getSquare(x, y int) (*square, error) {
	if x < 0 || x >= g.size || y < 0 || y >= g.size {
		return nil, errors.New("invalid coordinates: out of bounds")
	}

	return &g.minefield[y][x], nil
}

func (g *Game) explode() {
	g.gamestate = lost

	for _, row := range g.minefield {
		for _, s := range row {
			if s.mined {
				s.trigger()
			}
		}
	}
}

func (g *Game) Hit(x int, y int) error {
	if g.gamestate != playing {
		return errors.New("game over: start a new game to play again")
	}

	s, err := g.getSquare(x, y)

	if err != nil {
		return err
	}

	if s.triggered {
		return errors.New("cell already triggered")
	}

	s.trigger()

	if s.mined {
		g.explode()

		return nil
	}

	if s.value <= 0 {
		neighbors := s.getNeighbors(g.size)

		for _, pos := range *neighbors {
			if !g.minefield[pos.y][pos.x].triggered && !g.minefield[pos.y][pos.x].mined {
				g.Hit(pos.x, pos.y)
			}
		}
	}

	if g.movesLeft <= 0 {
		g.gamestate = won
	}

	return nil
}

func (g *Game) Flag(x, y int) error {
	if g.gamestate != playing {
		return errors.New("game over: start a new game to play again")
	}

	s, err := g.getSquare(x, y)

	if err != nil {
		return err
	}

	if s.triggered {
		return errors.New("cell already triggered")
	}

	s.flag()

	return nil
}
