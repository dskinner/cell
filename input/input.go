package input

import (
	"dasa.cc/cell/automata"
	"github.com/jteeuwen/glfw"
)

var Sim *automata.Life

func Init() {
	glfw.SetMouseWheelCallback(mouseWheelHandler)
	glfw.SetMouseButtonCallback(mouseButtonHandler)
	glfw.SetMousePosCallback(mousePosHandler)
	glfw.SetKeyCallback(keyHandler)
}
