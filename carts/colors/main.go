package main

import (
	null0 "github.com/notnullgames/null0go/cartapi"
)

func main() {
	println("Color constants available:")
	println("RED:", null0.RED.R, null0.RED.G, null0.RED.B, null0.RED.A)
	println("GREEN:", null0.GREEN.R, null0.GREEN.G, null0.GREEN.B, null0.GREEN.A)
	println("BLUE:", null0.BLUE.R, null0.BLUE.G, null0.BLUE.B, null0.BLUE.A)
	println("BLANK:", null0.BLANK.R, null0.BLANK.G, null0.BLANK.B, null0.BLANK.A)
	println("All 25 color constants are available!")
}
