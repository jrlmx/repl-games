package minesweeper

type square struct {
	pos
	value                     int
	mined, flagged, triggered bool
	neighbors                 *[]pos
}

func newSquare(x, y int) square {
	return square{
		pos:       pos{x: x, y: y},
		value:     0,
		mined:     false,
		flagged:   false,
		triggered: false,
	}
}

func (s *square) trigger() {
	s.triggered = true
}

func (s *square) flag() {
	s.flagged = !s.flagged
}

func (s *square) IsMine() bool {
	return s.mined
}

func (s *square) IsFlagged() bool {
	return s.flagged
}

func (s *square) IsTriggered() bool {
	return s.triggered
}

func (s *square) getValue() int {
	return s.value
}

func (s *square) getNeighbors(size int) *[]pos {
	if s.neighbors != nil {
		return s.neighbors
	}

	n := make([]pos, 3, 8)

	for y := max(s.y-1, 0); y <= min(size-1, s.y+1); y++ {
		for x := max(s.x-1, 0); x <= min(size-1, s.x+1); x++ {
			if x == s.x && y == s.y {
				continue
			}

			n = append(n, pos{x: x, y: y})
		}
	}

	s.neighbors = &n

	return s.neighbors
}
