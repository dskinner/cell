package automata

import (
	"errors"
	"fmt"
)

func NewRuler(name string) (Ruler, error) {
	switch name {
	case "conway":
		return new(Conway), nil
	case "wolfram":
		return new(Wolfram), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown rule: %s", name))
	}

	panic("unreachable")
}

// Rulers returns a list of all rules available, for listing from cli and
// eventually through the GUI.
func Rulers() []string {
	return []string {
		"conway",
		"wolfram",
	}
}

// State intends to be embedded in a Ruler implementation to provide generic
// information on a cell and its eight possible to determine if the cell is
// dead or alive.
type State struct{
	count int
	left, right, top, bottom, topLeft, topRight, bottomLeft, bottomRight bool
}

// SetState records which surrounding neighbors are dead and alive and keeps
// a count of all alive neighbors.
func (s *State) SetState(l *Life, i int) {
	// reset state
	s.count = 0
	s.left = false
	s.right = false
	s.top = false
	s.bottom = false
	s.topLeft = false
	s.topRight = false
	s.bottomLeft = false
	s.bottomRight = false

	// Each cell has eight neighbors. For cells along the grid boundary,
	// neighbors are considered by wrapping the grid

	// left
	if l.Cells[l.BoundValue(i-1)] {
		s.count++
		s.left = true
	}

	// right
	if l.Cells[l.BoundValue(i+1)] {
		s.count++
		s.right = true
	}

	// top
	if l.Cells[l.BoundValue(i-l.Size)] {
		s.count++
		s.top = true
	}

	// bottom
	if l.Cells[l.BoundValue(i+l.Size)] {
		s.count++
		s.bottom = true
	}

	// top left
	if l.Cells[l.BoundValue(i-(l.Size-1))] {
		s.count++
		s.topLeft = true
	}

	// top right
	if l.Cells[l.BoundValue(i-(l.Size+1))] {
		s.count++
		s.topRight = true
	}

	// bottom left
	if l.Cells[l.BoundValue(i+(l.Size-1))] {
		s.count++
		s.bottomLeft = true
	}

	// bottom right
	if l.Cells[l.BoundValue(i+(l.Size+1))] {
		s.count++
		s.bottomRight = true
	}
}

// Ruler defines how the game is played.
type Ruler interface {
	Rule(*Life, int) bool
}

// Conway is a standard implementation of Conway's Game of Life.
type Conway struct {
	State
}

func (c *Conway) Rule(l *Life, i int) bool {
	c.SetState(l, i)

	// the rules of life
	if l.Cells[i] && (c.count == 2 || c.count == 3) {
		return true
	}
	if !l.Cells[i] && c.count == 3 {
		return true
	}
	return false
}

// TODO rename to which rule number this is
type Wolfram struct {
	State
}

func (w *Wolfram) Rule(l *Life, i int) bool {
	w.SetState(l, i)

	if w.topLeft && w.top && w.topRight {
		return false
	}
	if w.topLeft && w.top && !w.topRight {
		return false
	}
	if w.topLeft && !w.top && w.topRight {
		return false
	}
	if w.topLeft && !w.top && !w.topRight {
		return true
	}
	if !w.topLeft && w.top && w.topRight {
		return false
	}
	if !w.topLeft && w.top && !w.topRight {
		return false
	}
	if !w.topLeft && !w.top && w.topRight {
		return true
	}

	return l.Cells[i]
}


