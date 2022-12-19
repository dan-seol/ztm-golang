package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/message"
)

func main() {
	message.Hi()
	display.Display("Hello from display")
	message.Exciting("an exciting message")
}
