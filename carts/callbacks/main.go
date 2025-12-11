package main

import "unsafe"

// Imported from host
//go:wasmimport host trace
func hostTrace(ptr uint32, len uint32)

func trace(s string) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	hostTrace(uint32(uintptr(ptr)), uint32(len(s)))
}

// main is required by TinyGo but we won't call _start from the host
func main() {
}

//export load
func load() {
	trace("[init] Callbacks example cart loaded!")
}

//export buttonUp
func buttonUp(button uint32, player uint32) {
	trace("[buttonUp] called!")
}

//export buttonDown
func buttonDown(button uint32, player uint32) {
	trace("[buttonDown] called!")
}

//export keyUp
func keyUp(key uint32) {
	trace("[keyUp] called!")
}

//export keyDown
func keyDown(key uint32) {
	trace("[keyDown] called!")
}

//export mouseDown
func mouseDown(button uint32) {
	trace("[mouseDown] called!")
}

//export mouseUp
func mouseUp(button uint32) {
	trace("[mouseUp] called!")
}

//export mouseMoved
func mouseMoved(x float32, y float32) {
	trace("[mouseMoved] called!")
}
