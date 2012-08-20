package main

import (
	"flag"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"log"
	"runtime"

	"dasa.cc/cell/input"
	"dasa.cc/cell/automata"
)

var flagSize = flag.Int("size", 640, "width and height of cell grid")
var flagSeed = flag.Int64("seed", 1, "value given for PRNG")
var flagDelay = flag.String("delay", "33ms", "value given in milliseconds to pause between steps")

const (
	Title  = "Universe"
	Width  = 640
	Height = 640
)

var life *automata.Life

func main() {
	runtime.GOMAXPROCS(2)
	flag.Parse()

	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 0, 0, glfw.Windowed); err != nil {
		log.Fatal(err)
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	initScene()
	defer destroyScene()

	for glfw.WindowParam(glfw.Opened) == 1 {
		drawScene()
		glfw.SwapBuffers()
	}
}

func initScene() {
	input.Init()
	life = automata.NewLife(*flagSize, *flagDelay, *flagSeed)
	input.Sim = life
	go life.Run()
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Viewport(0, 0, Width, Height)
}

func destroyScene() {

}

func drawCells() {

	if input.Zoom > 0 {
		gl.PointSize(gl.Float(float64(Width) / float64(*flagSize) * input.Zoom))
	}

	gl.Color4f(1, 1, 1, 1)
	gl.Enable(gl.POINT_SMOOTH)
	gl.Begin(gl.POINTS)

	for i, j := 0, 0; i < (life.Size * life.Size); i, j = i+life.Size, j+1 {
		for k, v := range life.Cells[i : i+life.Size] {
			if v {
				gl.Vertex2f(gl.Float(k), gl.Float(j))
			}
		}
	}

	gl.End()
	gl.Disable(gl.POINT_SMOOTH)
}

func drawGrid() {
	gl.Color4f(0.2, 0.2, 0.2, 1)
	gl.Begin(gl.LINES)
	for i := -0.5; i < Width; i += 1 {
		gl.Vertex2f(gl.Float(-0.5), gl.Float(i))
		gl.Vertex2f(gl.Float(life.Size), gl.Float(i))

		gl.Vertex2f(gl.Float(i), -0.5)
		gl.Vertex2f(gl.Float(i), gl.Float(life.Size))
	}
	gl.End()
}

var frameCount float64
var t0 = glfw.Time()

func fps() {
	t1 := glfw.Time()
	if t1-t0 >= 1.0 {
		log.Print(frameCount / (t1 - t0))
		log.Print(life.Generation)
		t0 = t1
		frameCount = 0
	} else {
		frameCount++
	}
}

func drawScene() {
	fps()
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	if input.Zoom > 0 {
		gl.Ortho(0, gl.Double(float64(*flagSize)/input.Zoom), gl.Double(float64(*flagSize)/input.Zoom), 0, -1, 1.0)
	}

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(gl.Float(input.TransX*1), gl.Float(input.TransY*1), 0)
	drawGrid()
	drawCells()
}
