package main

import (
	null0 "github.com/notnullgames/null0go/cartapi"
)

// called to set things up
func main() {
	println("[init] Callbacks example cart loaded!")
	println("Available types: Color, Vector, Rectangle, Dimensions")
	println("Available enums: Key, GamepadButton, MouseButton, ImageFilter, SfxPresetType")
}

//export buttonUp
func buttonUp(button uint32, player uint32) {
	println("[buttonUp] called!")
}

//export buttonDown
func buttonDown(button uint32, player uint32) {
	println("[buttonDown] called!")
}

//export keyUp
func keyUp(key uint32) {
	println("[keyUp] called!")
	// Example: check if it's the space key
	if key == uint32(null0.KEY_SPACE) {
		println("[keyUp] Space key!")
	}
}

//export keyDown
func keyDown(key uint32) {
	println("[keyDown] called!")
	// Example: check for specific keys
	if key == uint32(null0.KEY_A) {
		println("[keyDown] A key!")
	} else if key == uint32(null0.KEY_ESCAPE) {
		println("[keyDown] Escape key!")
	}
}

//export mouseDown
func mouseDown(button uint32) {
	println("[mouseDown] called!")
	// Example: check button type
	if button == uint32(null0.MOUSE_BUTTON_LEFT) {
		println("[mouseDown] Left click!")
	}
}

//export mouseUp
func mouseUp(button uint32) {
	println("[mouseUp] called!")
}

//export mouseMoved
func mouseMoved(x float32, y float32) {
	println("[mouseMoved] called!")
}
