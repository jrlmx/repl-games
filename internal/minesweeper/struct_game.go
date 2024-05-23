package minesweeper

import (
	"errors"
	"fmt"
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
			row[y] = newSquare(x, y)
		}

		minefield = append(minefield, row)
	}

	m := mines

	for m > 0 {
		for _, row := range minefield {
			for _, s := range row {
				if m == 0 {
					break
				}

				if !s.mined {
					s.mined = true
					m--

					for _, n := range *s.getNeighbors(size) {
						minefield[n.y][n.x].value++
					}
				}
			}
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
	for _, row := range g.minefield {
		for _, s := range row {
			if s.mined {
				s.trigger()
			}
		}
	}
}

func (g *Game) Hit(x, y int) error {
	s, err := g.getSquare(x, y)

	if err != nil {
		return err
	}

	if s.triggered {
		return fmt.Errorf("already hit (%v, %v)", x, y)
	}

	if s.mined {
		g.gamestate = lost
		g.explode()

		return nil
	}

	s.trigger()
	g.movesLeft--

	for _, n := range *s.getNeighbors(g.size) {
		ns, _ := g.getSquare(n.x, n.y)

		if ns.value == 0 && !ns.triggered {
			g.Hit(n.x, n.y)
		}
	}

	if g.movesLeft == 0 {
		g.gamestate = won
	}

	return nil
}

func (g *Game) Flag(x, y int) error {
	s, err := g.getSquare(x, y)

	if err != nil {
		return err
	}

	s.flag()

	return nil
}
