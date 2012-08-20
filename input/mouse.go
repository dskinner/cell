package input

import (
	"github.com/jteeuwen/glfw"
	"log"
)

var Zoom float64 = 1

var curButton int
var curButtonState int

var TransX float64
var TransY float64

var prevX, prevY, prevZoom int

func mouseWheelHandler(delta int) {

	Zoom += float64(delta-prevZoom) * 0.1
	prevZoom = delta

	if Zoom < 1 {
		Zoom = 1
	}
}

func mouseButtonHandler(button int, state int) {

	prevX, prevY = glfw.MousePos()
	curButton = button
	curButtonState = state

	if button == glfw.MouseRight && state == 1 {
		stickyX = prevX
		stickyY = prevY
		drawAt(prevX, prevY)
	}
}

func mousePosHandler(mx int, my int) {
	if curButton == 0 && curButtonState == 1 {

		// target size is 640, so adjust accordingly for panning speed
		TransX += float64(mx - prevX)/Zoom*float64(Sim.Size)/640
		TransY += float64(my - prevY)/Zoom*float64(Sim.Size)/640

		prevX, prevY = mx, my
	}

	if curButton == glfw.MouseRight && curButtonState == 1 {
		drawAt(mx, my)
	}

	// easier to draw on touchpad by holding alt and moving
	if ModAlt {
		drawAt(mx, my)
	}
}

var (
	stickyX int
	stickyY int
)

func drawAt(mx, my int) {
	if !ModShift {
		stickyY = my
	} else {
		my = stickyY
	}

	if !ModCtrl {
		stickyX = mx
	} else {
		mx = stickyX
	}

	x, y := float64(mx)-0.5, float64(my)-0.5
	s := float64(Sim.Size)
	z := Zoom
	dx, dy := TransX-0.5, TransY-0.5
	w := float64(640)

	row := ((y/z)-dy)*s/w
	col := ((x/z)-dx)*s/w

	log.Print(row, col)
	pos := Sim.Pos(int(row), int(col))
	if pos >= 0 && pos < len(Sim.Cells) {
		Sim.Cells[pos] = true
	}
}


