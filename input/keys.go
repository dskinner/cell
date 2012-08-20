package input

import (
	"github.com/jteeuwen/glfw"
)

var ModShift bool
var ModAlt bool
var ModCtrl bool

func keyHandler(key int, state int) {

	switch key {
	case glfw.KeyLshift, glfw.KeyRshift:
		ModShift = (state == 1)
	case glfw.KeyLalt, glfw.KeyRalt:
		ModAlt = (state == 1)
	case glfw.KeyLctrl, glfw.KeyRctrl:
		ModCtrl = (state == 1)
	case glfw.KeySpace:
		if state == 0 {
			break
		}
		if Sim.Running {
			Sim.Stop()
		} else {
			go Sim.Run()
		}
	case glfw.KeyRight:
		if state == 0 {
			break
		}
		if !Sim.Running {
			Sim.Step()
		}
	}

}
