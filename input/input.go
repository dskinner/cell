package input

import (
	"github.com/jteeuwen/glfw"
	"dasa.cc/cell/automata"
)

var Sim *automata.Life

func Init() {
	glfw.SetMouseWheelCallback(mouseWheelHandler)
	glfw.SetMouseButtonCallback(mouseButtonHandler)
	glfw.SetMousePosCallback(mousePosHandler)
	glfw.SetKeyCallback(keyHandler)
}
