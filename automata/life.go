package automata

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Life struct {
	Cells      []bool
	Future     []bool
	Size       int
	Delay      time.Duration
	Generation float64
	Running    bool
}

func NewLife(size int, delay string, seed int64) (l *Life) {
	l = new(Life)
	l.Cells = make([]bool, size*size)
	l.Future = make([]bool, size*size)
	l.Size = size

	d, err := time.ParseDuration(delay)
	if err != nil {
		log.Fatal(err)
	}
	l.Delay = d

	// init random state
	if seed != 0 {
		rand.Seed(seed)
		for i := range l.Cells {
			l.Cells[i] = (rand.Int() % 2) == 0
		}
	} else {
		i := size / 2
		l.Cells[i-1] = true
		l.Cells[i] = true
		l.Cells[i+1] = true
	}

	return l
}

func (l *Life) Pos(row, col int) int {
	return (row * l.Size) + col
}

func (l *Life) BoundValue(i int) int {
	if i < 0 {
		return len(l.Cells) + i
	}

	if i >= len(l.Cells) {
		return i - len(l.Cells)
	}

	return i
}

func (l *Life) MakeObject(name string) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	/*
		for _, coord := range coords {
			l.Cells[l.Pos(coord[0], coord[1])] = true
		}
	*/
}

func (l *Life) Step() {
	// determine the future
	for i := range l.Cells {
		l.Future[i] = l.IsAlive(i)
	}

	// swap the pointers
	l.Cells, l.Future = l.Future, l.Cells
	l.Generation++
}

func (l *Life) Done() bool {
	for i, v := range l.Cells {
		if l.Future[i] != v {
			return false
		}
	}
	return true
}

func (l *Life) IsAlive(i int) bool {
	count := 0
	//var t, tl, tr bool

	// Each cell has eight neighbors. For cells along the grid boundary,
	// neighbors are considered by wrapping the grid

	// left
	if l.Cells[l.BoundValue(i-1)] {
		count++
	}

	// right
	if l.Cells[l.BoundValue(i+1)] {
		count++
	}

	// top
	if l.Cells[l.BoundValue(i-l.Size)] {
		count++
		//t = true
	}

	// bottom
	if l.Cells[l.BoundValue(i+l.Size)] {
		count++
	}

	// top left
	if l.Cells[l.BoundValue(i-(l.Size-1))] {
		count++
		//tl = true
	}

	// top right
	if l.Cells[l.BoundValue(i-(l.Size+1))] {
		count++
		//tr = true
	}

	// bottom left
	if l.Cells[l.BoundValue(i+(l.Size-1))] {
		count++
	}

	// bottom right
	if l.Cells[l.BoundValue(i+(l.Size+1))] {
		count++
	}

	// the rules of life

	if l.Cells[i] && (count == 2 || count == 3) {
		return true
	}
	if !l.Cells[i] && count == 3 {
		return true
	}
	return false

	/*
		if tl && t && tr {
			return false
		}
		if tl && t && !tr {
			return false
		}
		if tl && !t && tr {
			return false
		}
		if tl && !t && !tr {
			return true
		}
		if !tl && t && tr {
			return false
		}
		if !tl && t && !tr {
			return false
		}
		if !tl && !t && tr {
			return true
		}

		return l.Cells[i]
	*/
}

func (l *Life) Run() {
	l.Running = true
	for {
		if !l.Running {
			break
		}
		t := time.Now()
		l.Step()
		/*
			if l.Done() {
				break
			}
		*/
		time.Sleep(l.Delay - time.Now().Sub(t))
	}
}

func (l *Life) Stop() {
	l.Running = false
}

func (l *Life) Print() {
	var buf bytes.Buffer

	for i := 0; i < (l.Size * l.Size); i += l.Size {
		for _, v := range l.Cells[i : i+l.Size] {
			if v {
				buf.WriteRune('0')
			} else {
				buf.WriteRune(' ')
			}
		}
		buf.WriteRune('\n')
	}

	fmt.Println(buf.String())
}
